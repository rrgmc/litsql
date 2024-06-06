package structargs

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rrgmc/litsql"
)

// New returns a [litsql.ArgValues] from struct fields.
func New(v any, options ...Option) (litsql.ArgValues, error) {
	optns := argValuesOptions{
		tagName: "json",
	}
	for _, opt := range options {
		opt(&optns)
	}
	return NewFromConfig(&mapstructure.DecoderConfig{
		TagName: optns.tagName,
	}, v)
}

// NewFromConfig returns a [litsql.ArgValues] from struct fields using a [mapstructure.DecoderConfig].
func NewFromConfig(config *mapstructure.DecoderConfig, v any) (litsql.ArgValues, error) {
	result := map[string]any{}
	config.Result = &result

	dec, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, err
	}

	err = dec.Decode(v)
	if err != nil {
		return nil, err
	}

	return litsql.MapArgValues(result), nil
}

type Option func(*argValuesOptions)

type argValuesOptions struct {
	tagName string
}

// WithTagName sets the struct tag name to use. Default is "json".
func WithTagName(tagName string) Option {
	return func(o *argValuesOptions) {
		o.tagName = tagName
	}
}
