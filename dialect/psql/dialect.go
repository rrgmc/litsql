package psql

import (
	"strconv"

	"github.com/rrgmc/litsql"
)

//nolint:gochecknoglobals
var (
	Dialect     dialect
	dollar      = "$"
	doubleQuote = `"`
)

type dialect struct{}

func (d dialect) WriteArg(w litsql.Writer, position int) {
	w.Write(dollar)
	w.Write(strconv.Itoa(position))
}

func (d dialect) WriteQuoted(w litsql.Writer, s string) {
	w.Write(doubleQuote)
	w.Write(s)
	w.Write(doubleQuote)
}
