package expr

import (
	"github.com/rrgmc/litsql"
)

// Raw is an [litsql.Expression] that outputs the string itself as-is.
type Raw string

func (r Raw) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	w.Write(string(r))
	return nil, nil
}
