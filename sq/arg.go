package sq

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal"
)

// Arg adds a named argument.
func Arg(name string) litsql.Argument {
	return &internal.NamedArgument{ArgName: name}
}

// ArgDefault adds a named argument with a default value.
func ArgDefault(name string, defaultValue any) litsql.Argument {
	return &internal.NamedArgumentWithDefault{ArgName: name, DefaultValue: defaultValue}
}

// DBArg adds a DB named argument.
func DBArg(name string) litsql.Argument {
	return &internal.DBNamedArgument{ArgName: name}
}

// DBArgDefault adds a DB named argument with a default value.
func DBArgDefault(name string, defaultValue any) litsql.Argument {
	return &internal.DBNamedArgumentWithDefault{ArgName: name, DefaultValue: defaultValue}
}

// ArgFunc returns the argument value in a callback.
func ArgFunc(fn func() (any, error)) litsql.Argument {
	return &internal.FuncArgument{FN: fn}
}

// ParseArgs replaces all [litsql.Argument] instances in args with named values.
func ParseArgs(args []any, values ...any) ([]any, error) {
	return internal.ParseArgs(args, values...)
}

// ParseArgValues replaces all [litsql.Argument] instances in args with named values.
func ParseArgValues(args []any, values ...litsql.ArgValues) ([]any, error) {
	return internal.ParseArgValues(args, values...)
}
