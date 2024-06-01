package sqlite

import (
	"strconv"
	"strings"

	"github.com/rrgmc/litsql"
)

//nolint:gochecknoglobals
var (
	Dialect      dialect
	questionMark = "?"
	colon        = ":"
	doubleQuote  = `"`
)

type dialect struct{}

func (d dialect) WriteArg(w litsql.Writer, position int) {
	w.Write(questionMark)
	w.Write(strconv.Itoa(position))
}

func (d dialect) WriteNamedArg(w litsql.Writer, name string) {
	w.Write(colon)
	w.Write(name)
}

func (d dialect) WriteQuoted(w litsql.Writer, s string) {
	w.Write(doubleQuote)
	w.Write(s)
	w.Write(doubleQuote)
}

func (d dialect) WriteCheckQuoted(w litsql.Writer, s string) {
	if !strings.ContainsAny(s, " ") {
		w.Write(s)
		return
	}
	d.WriteQuoted(w, s)
}
