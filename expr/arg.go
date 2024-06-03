package expr

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal"
)

// Arg outputs the dialect-specific argument matcher, and wraps the passed value.
func Arg(value any) litsql.Expression {
	return arg{value: value}
}

// Args wraps multiple values in Arg.
func Args(values []any) []litsql.Expression {
	var ret []litsql.Expression
	for _, v := range values {
		ret = append(ret, Arg(v))
	}
	return ret
}

// In outputs the list of values as Arg separated by commas.
func In(values ...any) litsql.Expression {
	return argList{
		values:    values,
		separator: internal.CommaSpace,
	}
}

// InP outputs the list of values as Arg separated by commas, wrapped in parentheses.
func InP(values ...any) litsql.Expression {
	return argList{
		values:    values,
		separator: internal.CommaSpace,
		prefix:    internal.OpenPar,
		suffix:    internal.ClosePar,
	}
}

type arg struct {
	value any
}

func (a arg) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) (args []any, err error) {
	return clauseWriteArg(w, d, start, a.value)
}

type argList struct {
	values    []any
	separator string
	prefix    string
	suffix    string
}

func (a argList) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) (args []any, err error) {
	w.Write(a.prefix)

	startAt := start
	for i, v := range a.values {
		if i > 0 {
			w.Write(a.separator)
		}
		newArgs, err := clauseWriteArg(w, d, startAt, v)
		if err != nil {
			return nil, err
		}
		startAt += len(newArgs)
		args = append(args, newArgs...)
	}

	w.Write(a.suffix)

	return args, nil
}
