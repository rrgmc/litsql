package reflectxargs

import (
	"reflect"

	"github.com/iancoleman/strcase"
	"github.com/jmoiron/sqlx/reflectx"
	"github.com/rrgmc/litsql"
)

// New returns a [litsql.ArgValues] from struct fields. If value is not a struct, returns nil.
func New(value any, options ...Option) litsql.ArgValues {
	v := getReflectValue(value)
	if !v.IsValid() {
		return nil
	}
	optns := argValuesOptions{
		tagName: "json",
	}
	for _, opt := range options {
		opt(&optns)
	}
	if optns.mapper == nil {
		if optns.mapperFunc == nil {
			optns.mapperFunc = strcase.ToSnake
		}
		optns.mapper = reflectx.NewMapperFunc(optns.tagName, optns.mapperFunc)
	}
	return newFromMapper(v, optns.mapper)
}

// NewFromMapper returns a [litsql.ArgValues] from struct fields. If value is not a struct, returns nil.
func NewFromMapper(value any, mapper *reflectx.Mapper) litsql.ArgValues {
	v := getReflectValue(value)
	if !v.IsValid() {
		return nil
	}
	return newFromMapper(v, mapper)
}

func newFromMapper(value reflect.Value, mapper *reflectx.Mapper) litsql.ArgValues {
	return &argValues{
		sm:    mapper.TypeMap(value.Type()),
		value: value,
	}
}

type argValues struct {
	sm    *reflectx.StructMap
	value reflect.Value
}

func (s *argValues) Get(name string) (any, bool) {
	f, ok := s.sm.Names[name]
	if !ok {
		return nil, false
	}
	return reflectx.FieldByIndexes(s.value, f.Index).Interface(), true
}

type Option func(*argValuesOptions)

type argValuesOptions struct {
	mapper     *reflectx.Mapper
	tagName    string
	mapperFunc func(string) string
}

// WithTagName sets the struct tag name to use. Default is "json".
func WithTagName(tagName string) Option {
	return func(o *argValuesOptions) {
		o.tagName = tagName
	}
}

// WithMapperFunc sets the field name mapper function.
func WithMapperFunc(mapperFunc func(string) string) Option {
	return func(o *argValuesOptions) {
		o.mapperFunc = mapperFunc
	}
}

// WithMapper sets the [reflectx.Mapper] instance. If set, the other options will be ignored
func WithMapper(mapper *reflectx.Mapper) Option {
	return func(o *argValuesOptions) {
		o.mapper = mapper
	}
}
