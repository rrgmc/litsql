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

// WithGetArgValuesInstanceOptions adds query parse args options.
func WithGetArgValuesInstanceOptions(options ...GetArgValuesInstanceOption) BuildOption {
	return internal.WithBuildQueryGetArgValuesInstanceOptions(options...)
}

// WithParseArgs adds named argument values.
func WithParseArgs(argValues any) BuildOption {
	return internal.WithBuildQueryParseArgs(argValues)
}

// WithParseArgValues adds named argument values.
func WithParseArgValues(argValues litsql.ArgValues) BuildOption {
	return internal.WithBuildQueryParseArgValues(argValues)
}
