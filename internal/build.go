package internal

import (
	"bytes"

	"github.com/rrgmc/litsql"
)

// BuildQuery builds a query string and its arguments.
func BuildQuery(q litsql.Query, options ...BuildQueryOption) (string, []any, error) {
	var optns buildQueryOptions
	for _, opt := range options {
		opt(&optns)
	}

	var b bytes.Buffer
	w := NewWriter(&b, optns.writerOptions...)
	args, err := q.WriteQuery(w, 1)
	if err != nil {
		return "", nil, err
	}
	if w.Err() != nil {
		return "", nil, err
	}

	if len(optns.parseArgs) > 0 || len(optns.parseArgValues) > 0 {
		if len(optns.parseArgs) > 0 {
			args, err = ParseArgs(args, append(optns.parseArgs, ToAnySlice(optns.parseArgValues)...)...)
		} else {
			args, err = ParseArgValues(args, optns.parseArgValues...)
		}
		if err != nil {
			return "", nil, err
		}
	}

	return b.String(), args, nil
}

type BuildQueryOption func(options *buildQueryOptions)

type buildQueryOptions struct {
	writerOptions  []WriterOption
	parseArgs      []any
	parseArgValues []litsql.ArgValues
}

// WithBuildQueryWriterOptions adds writer options.
func WithBuildQueryWriterOptions(writerOptions ...WriterOption) BuildQueryOption {
	return func(options *buildQueryOptions) {
		options.writerOptions = append(options.writerOptions, writerOptions...)
	}
}

// WithBuildQueryParseArgs adds named argument values.
func WithBuildQueryParseArgs(argValues ...any) BuildQueryOption {
	return func(options *buildQueryOptions) {
		options.parseArgs = append(options.parseArgs, argValues...)
	}
}

// WithBuildQueryParseArgValues adds named argument values.
func WithBuildQueryParseArgValues(argValues ...litsql.ArgValues) BuildQueryOption {
	return func(options *buildQueryOptions) {
		options.parseArgValues = append(options.parseArgValues, argValues...)
	}
}
