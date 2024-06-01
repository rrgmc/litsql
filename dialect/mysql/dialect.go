package mysql

import (
	"strings"

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

func (d dialect) WriteCheckQuoted(w litsql.Writer, s string) {
	if !strings.ContainsAny(s, " ") {
		w.Write(s)
		return
	}
	d.WriteQuoted(w, s)
}
