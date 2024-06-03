package sq

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal"
)

// Build builds a query string and its arguments.
func Build(q litsql.Query, writerOptions ...WriterOption) (string, Args, error) {
	return internal.BuildQuery(q, writerOptions...)
}
