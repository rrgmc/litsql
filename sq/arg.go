package sq

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal"
)

// NamedArg adds a named argument for replacement with ParseArgs.
func NamedArg(name string, options ...ArgOption) litsql.Argument {
	var optns argOptions
	for _, opt := range options {
		opt(&optns)
	}
	if optns.defaultValue != nil {
		return &internal.NamedArgumentWithDefault{ArgName: name, DefaultValue: *optns.defaultValue}
	}
	return &internal.NamedArgument{ArgName: name}
}

// DBNamedArg adds a DB-native named argument (for databases that support it).
// The end value will be wrapped in [sql.Named].
func DBNamedArg(name string, options ...ArgOption) litsql.Argument {
	var optns argOptions
	for _, opt := range options {
		opt(&optns)
	}
	if optns.defaultValue != nil {
		return &internal.DBNamedArgumentWithDefault{ArgName: name, DefaultValue: *optns.defaultValue}
	}
	return &internal.DBNamedArgument{ArgName: name}
}

// ArgFunc returns the argument value in a callback.
func ArgFunc(fn func() (any, error)) litsql.Argument {
	return &internal.FuncArgument{FN: fn}
}

// ParseArgs replaces all [litsql.Argument] instances in args with named values.
func ParseArgs(args []any, values any, options ...GetArgValuesInstanceOption) ([]any, error) {
	return internal.ParseArgs(args, values, options...)
}

// ParseArgValues replaces all [litsql.Argument] instances in args with named values.
func ParseArgValues(args []any, values litsql.ArgValues) ([]any, error) {
	return internal.ParseArgValues(args, values)
}

type GetArgValuesInstanceOption = internal.GetArgValuesInstanceOption

// GetArgValuesInstance gets the [litsql.ArgValues] instance from the passed parameter.
func GetArgValuesInstance(values any, options ...GetArgValuesInstanceOption) (litsql.ArgValues, error) {
	return internal.GetArgValuesInstance(values, options...)
}

// WithGetArgValuesInstanceOptionCustom sets a custom [litsql.ArgValues] detector.
func WithGetArgValuesInstanceOptionCustom(custom func(values any) (litsql.ArgValues, error)) GetArgValuesInstanceOption {
	return internal.WithGetArgValuesInstanceOptionCustom(custom)
}

// ArgsParser wraps parseable argument results.
type ArgsParser []any

func (a ArgsParser) Parse(values any) ([]any, error) {
	return internal.ParseArgs(a, values)
}

func (a ArgsParser) ParseValues(values litsql.ArgValues) ([]any, error) {
	return internal.ParseArgValues(a, values)
}

type ArgOption func(options *argOptions)

// WithDefaultValue sets a default value if the argument name was not passed.
func WithDefaultValue(defaultValue any) ArgOption {
	return func(options *argOptions) {
		options.defaultValue = &defaultValue
	}
}

type argOptions struct {
	defaultValue *any
}
