package mysql

import (
	"github.com/rrgmc/litsql"
)

//nolint:gochecknoglobals
var (
	Dialect      dialect
	questionMark = "?"
	backtick     = "`"
)

type dialect struct{}

func (d dialect) WriteArg(w litsql.Writer, position int) {
	w.Write(questionMark)
}

func (d dialect) WriteQuoted(w litsql.Writer, s string) {
	w.Write(backtick)
	w.Write(s)
	w.Write(backtick)
}
