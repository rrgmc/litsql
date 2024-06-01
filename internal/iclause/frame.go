package iclause

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
)

type Frame struct {
	Defined   bool // whether any of the parts was defined
	Mode      string
	Start     litsql.Expression
	End       litsql.Expression // can be nil
	Exclusion string            // can be empty
}

func (f *Frame) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	b := litsql.NewExpressBuilder(w, d, start)

	if f.Mode == "" {
		f.Mode = "RANGE"
	}
	if f.Start == nil {
		f.Start = expr.Raw("UNBOUNDED PRECEDING")
	}

	w.Write(f.Mode)
	w.Write(" ")
	if f.End != nil {
		w.Write("BETWEEN ")
	}
	b.Express(f.Start)
	b.ExpressIf(f.End, f.End != nil, expr.Raw(" AND "), nil)

	b.ExpressIf(expr.Raw(f.Exclusion), f.Exclusion != "", expr.Raw(" EXCLUDE "), nil)

	return b.Result()
}

func (f *Frame) SetMode(mode string) {
	f.Defined = true
	f.Mode = mode
}

func (f *Frame) SetStart(start litsql.Expression) {
	f.Defined = true
	f.Start = start
}

func (f *Frame) SetEnd(end litsql.Expression) {
	f.Defined = true
	f.End = end
}

func (f *Frame) SetExclusion(excl string) {
	f.Defined = true
	f.Exclusion = excl
}
