package sq

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal"
)

type BuildQueryOption = internal.BuildQueryOption

// Build builds a query string and its arguments.
func Build(q litsql.Query, options ...BuildQueryOption) (string, []any, error) {
	return internal.BuildQuery(q, options...)
}

// WithWriterOptions adds writer options.
func WithWriterOptions(writerOptions ...WriterOption) BuildQueryOption {
	return internal.WithBuildQueryWriterOptions(writerOptions...)
}

// WithParseArgs adds named argument values.
func WithParseArgs(argValues any) BuildQueryOption {
	return internal.WithBuildQueryParseArgs(argValues)
}

// WithParseArgValues adds named argument values.
func WithParseArgValues(argValues litsql.ArgValues) BuildQueryOption {
	return internal.WithBuildQueryParseArgValues(argValues)
}
