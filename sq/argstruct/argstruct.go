package argstruct

import (
	"reflect"

	"github.com/rrgmc/litsql"
)

// Values returns a [litsql.ArgValues] from struct fields. If value is not a struct, returns nil.
func Values(value any, options ...Option) litsql.ArgValues {
	var optns argValuesOptions
	for _, opt := range options {
		opt(&optns)
	}
	return values(value, optns)
}

type ValuesFunc func(value any) litsql.ArgValues

// ValuesProvider returns a function that provides Values with fixed options.
func ValuesProvider(options ...Option) ValuesFunc {
	var optns argValuesOptions
	for _, opt := range options {
		opt(&optns)
	}
	return func(value any) litsql.ArgValues {
		return values(value, optns)
	}
}

func values(value any, options argValuesOptions) litsql.ArgValues {
	v, err := getReflectValue(value)
	if err != nil {
		return &errArgValues{err: err}
	}
	return valuesFromReflect(v, options)
}

func valuesFromReflect(value reflect.Value, options argValuesOptions) litsql.ArgValues {
	return &argValues{
		value:   value,
		options: options,
	}
}

type argValues struct {
	value   reflect.Value
	options argValuesOptions
}

func (s *argValues) Get(name string) (any, bool, error) {
	value, ok := s.getStructFieldByName(s.value, name)
	return value, ok, nil
}

type Option func(*argValuesOptions)

// WithTagName sets the struct tag name to use. Default is "json".
func WithTagName(tagName string) Option {
	return func(o *argValuesOptions) {
		o.tagName = tagName
	}
}

// WithDerefPointer dereferences pointers in struct field values.
func WithDerefPointer(deref bool) Option {
	return func(o *argValuesOptions) {
		o.derefPointer = deref
	}
}

// WithMapperFunc sets the field name mapper function.
func WithMapperFunc(mapperFunc func(string) string) Option {
	return func(o *argValuesOptions) {
		o.mapperFunc = mapperFunc
	}
}

type argValuesOptions struct {
	tagName      string
	derefPointer bool
	mapperFunc   func(string) string
}

type errArgValues struct {
	err error
}

func (e *errArgValues) Get(s string) (any, bool, error) {
	return nil, false, e.err
}
