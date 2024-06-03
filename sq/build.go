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

// WithBuildWriterOptions adds writer options.
func WithBuildWriterOptions(writerOptions ...WriterOption) BuildQueryOption {
	return internal.WithBuildQueryWriterOptions(writerOptions...)
}

// WithBuildParseArgs adds named argument values.
func WithBuildParseArgs(argValues any) BuildQueryOption {
	return internal.WithBuildQueryParseArgs(argValues)
}

// WithBuildParseArgValues adds named argument values.
func WithBuildParseArgValues(argValues litsql.ArgValues) BuildQueryOption {
	return internal.WithBuildQueryParseArgValues(argValues)
}
