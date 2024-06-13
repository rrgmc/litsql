package sq

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal"
)

type BuildOption = internal.BuildQueryOption

// Build builds a query string and its arguments.
func Build(q litsql.Query, options ...BuildOption) (string, []any, error) {
	return internal.BuildQuery(q, options...)
}

// WithWriterOptions adds writer options.
func WithWriterOptions(writerOptions ...WriterOption) BuildOption {
	return internal.WithBuildQueryWriterOptions(writerOptions...)
}

// WithParseArgs adds named argument values.
func WithParseArgs(argValues litsql.ArgValues) BuildOption {
	return internal.WithBuildQueryParseArgs(argValues)
}
