package sq

import (
	"slices"

	"github.com/rrgmc/litsql"
)

// Handler bundles default parameters for functions.
type Handler interface {
	Build(q litsql.Query, options ...BuildOption) (string, []any, error)
	ParseArgs(args []any, values litsql.ArgValues) ([]any, error)
}

func NewHandler(options ...HandlerOption) Handler {
	ret := &handler{}
	for _, opt := range options {
		opt(ret)
	}
	return ret
}

type handler struct {
	buildOptions []BuildOption
}

func (h *handler) Build(q litsql.Query, options ...BuildOption) (string, []any, error) {
	return Build(q, slices.Concat(h.buildOptions, options)...)
}

func (h *handler) ParseArgs(args []any, values litsql.ArgValues) ([]any, error) {
	return ParseArgs(args, values)
}

type HandlerOption func(h *handler)

// WithDefaultBuildOptions sets default options for [Handler.Build].
func WithDefaultBuildOptions(options ...BuildOption) HandlerOption {
	return func(h *handler) {
		h.buildOptions = append(h.buildOptions, options...)
	}
}
