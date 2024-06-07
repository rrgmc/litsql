package argstruct

import (
	"reflect"

	"github.com/rrgmc/litsql"
)

// New returns a [litsql.ArgValues] from struct fields. If value is not a struct, returns nil.
func New(value any, options ...Option) litsql.ArgValues {
	v := getReflectValue(value)
	if !v.IsValid() {
		return nil
	}

	ret := &argValues{
		tagName: "json",
		value:   v,
	}
	for _, opt := range options {
		opt(ret)
	}
	return ret
}

type argValues struct {
	value reflect.Value

	tagName      string
	derefPointer bool
	mapperFunc   func(string) string
}

func (s *argValues) Get(name string) (any, bool) {
	return s.getStructFieldByName(s.value, name)
}

type Option func(*argValues)

// WithTagName sets the struct tag name to use. Default is "json".
func WithTagName(tagName string) Option {
	return func(o *argValues) {
		o.tagName = tagName
	}
}

// WithDerefPointer dereferences pointers in struct field values.
func WithDerefPointer(deref bool) Option {
	return func(o *argValues) {
		o.derefPointer = deref
	}
}

// WithMapperFunc sets the field name mapper function.
func WithMapperFunc(mapperFunc func(string) string) Option {
	return func(o *argValues) {
		o.mapperFunc = mapperFunc
	}
}
