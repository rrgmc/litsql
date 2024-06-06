package sq

import (
	"slices"

	"github.com/rrgmc/litsql"
)

// Handler bundles default parameters for functions.
type Handler interface {
	Build(q litsql.Query, options ...BuildOption) (string, []any, error)
	ParseArgs(args []any, values any, options ...GetArgValuesInstanceOption) ([]any, error)
	ParseArgValues(args []any, values litsql.ArgValues) ([]any, error)
	GetArgValuesInstance(values any, options ...GetArgValuesInstanceOption) (litsql.ArgValues, error)
}

func NewHandler(options ...HandlerOption) Handler {
	ret := &handler{}
	for _, opt := range options {
		opt(ret)
	}
	return ret
}

type handler struct {
	getArgValuesInstanceOption []GetArgValuesInstanceOption
	buildOptions               []BuildOption
}

func (h *handler) Build(q litsql.Query, options ...BuildOption) (string, []any, error) {
	return Build(q, slices.Concat(
		[]BuildOption{WithGetArgValuesInstanceOptions(h.getArgValuesInstanceOption...)},
		h.buildOptions,
		options,
	)...)
}

func (h *handler) ParseArgs(args []any, values any, options ...GetArgValuesInstanceOption) ([]any, error) {
	return ParseArgs(args, values, slices.Concat(h.getArgValuesInstanceOption, options)...)
}

func (h *handler) ParseArgValues(args []any, values litsql.ArgValues) ([]any, error) {
	return ParseArgValues(args, values)
}

func (h *handler) GetArgValuesInstance(values any, options ...GetArgValuesInstanceOption) (litsql.ArgValues, error) {
	return GetArgValuesInstance(values, slices.Concat(h.getArgValuesInstanceOption, options)...)
}

type HandlerOption func(h *handler)

// WithDefaultBuildOptions sets default options for [Handler.Build].
func WithDefaultBuildOptions(options ...BuildOption) HandlerOption {
	return func(h *handler) {
		h.buildOptions = append(h.buildOptions, options...)
	}
}

// WithDefaultGetArgValuesInstanceOptions sets default options for [Handler.ParseArgs] and [Handler.GetArgValuesInstance].
func WithDefaultGetArgValuesInstanceOptions(options ...GetArgValuesInstanceOption) HandlerOption {
	return func(h *handler) {
		h.getArgValuesInstanceOption = append(h.getArgValuesInstanceOption, options...)
	}
}
