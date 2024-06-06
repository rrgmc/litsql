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
	if optns.argValues == nil && optns.rawArgValues != nil {
		var err error
		optns.argValues, err = GetArgValuesInstance(optns.rawArgValues, optns.getArgValuesInstanceOption...)
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

type BuildQueryOption func(options *buildQueryOptions)

type buildQueryOptions struct {
	writerOptions              []WriterOption
	getArgValuesInstanceOption []GetArgValuesInstanceOption
	rawArgValues               any
	argValues                  litsql.ArgValues
}

// WithBuildQueryWriterOptions adds writer options.
func WithBuildQueryWriterOptions(writerOptions ...WriterOption) BuildQueryOption {
	return func(options *buildQueryOptions) {
		options.writerOptions = append(options.writerOptions, writerOptions...)
	}
}

// WithBuildQueryGetArgValuesInstanceOptions adds query parse args options.
func WithBuildQueryGetArgValuesInstanceOptions(options ...GetArgValuesInstanceOption) BuildQueryOption {
	return func(o *buildQueryOptions) {
		o.getArgValuesInstanceOption = append(o.getArgValuesInstanceOption, options...)
	}
}

// WithBuildQueryParseArgs adds named argument values.
func WithBuildQueryParseArgs(argValues any) BuildQueryOption {
	return func(o *buildQueryOptions) {
		o.rawArgValues = argValues
	}
}

// WithBuildQueryParseArgValues adds named argument values.
func WithBuildQueryParseArgValues(argValues litsql.ArgValues) BuildQueryOption {
	return func(options *buildQueryOptions) {
		options.argValues = argValues
	}
}
