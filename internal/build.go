package internal

import (
	"bytes"

	"github.com/rrgmc/litsql"
)

// BuildQuery builds a query string and its arguments.
func BuildQuery(q litsql.Query, options ...BuildQueryOption) (string, []any, error) {
	var optns buildQueryOptions
	for _, opt := range options {
		err := opt(&optns)
		if err != nil {
			return "", nil, err
		}
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

	if optns.argValues != nil {
		args, err = ParseArgValues(args, optns.argValues)
		if err != nil {
			return "", nil, err
		}
	}

	return b.String(), args, nil
}

type BuildQueryOption func(options *buildQueryOptions) error

type buildQueryOptions struct {
	writerOptions []WriterOption
	argValues     litsql.ArgValues
}

// WithBuildQueryWriterOptions adds writer options.
func WithBuildQueryWriterOptions(writerOptions ...WriterOption) BuildQueryOption {
	return func(options *buildQueryOptions) error {
		options.writerOptions = append(options.writerOptions, writerOptions...)
		return nil
	}
}

// WithBuildQueryParseArgs adds named argument values.
func WithBuildQueryParseArgs(argValues any) BuildQueryOption {
	return func(options *buildQueryOptions) error {
		av, err := GetArgValuesInstance(argValues)
		if err != nil {
			return err
		}
		options.argValues = av
		return nil
	}
}

// WithBuildQueryParseArgValues adds named argument values.
func WithBuildQueryParseArgValues(argValues litsql.ArgValues) BuildQueryOption {
	return func(options *buildQueryOptions) error {
		options.argValues = argValues
		return nil
	}
}
