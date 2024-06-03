package sq

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal"
)

type BuildQueryOption = internal.BuildQueryOption

// Build builds a query string and its arguments.
func Build(q litsql.Query, options ...BuildQueryOption) (string, Args, error) {
	return internal.BuildQuery(q, options...)
}

// WithBuildQueryWriterOptions adds writer options.
func WithBuildQueryWriterOptions(writerOptions ...WriterOption) BuildQueryOption {
	return internal.WithBuildQueryWriterOptions(writerOptions...)
}

// WithBuildQueryParseArgs adds named argument values.
func WithBuildQueryParseArgs(argValues ...any) BuildQueryOption {
	return internal.WithBuildQueryParseArgs(argValues...)
}

// WithBuildQueryParseArgValues adds named argument values.
func WithBuildQueryParseArgValues(argValues ...litsql.ArgValues) BuildQueryOption {
	return internal.WithBuildQueryParseArgValues(argValues...)
}
