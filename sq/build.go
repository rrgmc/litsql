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

// WithBuildQueryGetArgValuesInstanceOptions adds query parse args options.
func WithBuildQueryGetArgValuesInstanceOptions(options ...GetArgValuesInstanceOption) BuildQueryOption {
	return internal.WithBuildQueryGetArgValuesInstanceOptions(options...)
}

// WithParseArgs adds named argument values.
func WithParseArgs(argValues any) BuildQueryOption {
	return internal.WithBuildQueryParseArgs(argValues)
}

// WithParseArgValues adds named argument values.
func WithParseArgValues(argValues litsql.ArgValues) BuildQueryOption {
	return internal.WithBuildQueryParseArgValues(argValues)
}
