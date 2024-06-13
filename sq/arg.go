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
// Use [litsql.MapArgValues] to use a "map[string]any" as source.
func ParseArgs(args []any, values litsql.ArgValues) ([]any, error) {
	return internal.ParseArgs(args, values)
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
