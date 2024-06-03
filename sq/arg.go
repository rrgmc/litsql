package sq

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal"
)

// Arg adds a named argument.
func Arg(name string) litsql.Argument {
	return namedArgument{name: name}
}

// ArgDefault adds a named argument with a default value.
func ArgDefault(name string, defaultValue any) litsql.Argument {
	return namedArgumentWithDefault{name: name, defaultValue: defaultValue}
}

// DBArg adds a DB named argument.
func DBArg(name string) litsql.Argument {
	return dbNamedArgument{name: name}
}

// DBArgDefault adds a DB named argument with a default value.
func DBArgDefault(name string, defaultValue any) litsql.Argument {
	return dbNamedArgumentWithDefault{name: name, defaultValue: defaultValue}
}

// ArgFunc returns the argument value in a callback.
func ArgFunc(fn func() (any, error)) litsql.Argument {
	return funcArgument{fn: fn}
}

// Args wraps parseable argument results.
type Args []any

func (a Args) Parse(values ...any) ([]any, error) {
	return ParseArgs(a, values...)
}

func (a Args) ParseValues(values ...litsql.ArgValues) ([]any, error) {
	return ParseArgValues(a, values...)
}

// ParseArgs replaces all [litsql.Argument] instances in args with named values.
func ParseArgs(args []any, values ...any) ([]any, error) {
	return internal.ParseArgs(args, values...)
}

// ParseArgValues replaces all [litsql.Argument] instances in args with named values.
func ParseArgValues(args []any, values ...litsql.ArgValues) ([]any, error) {
	return internal.ParseArgValues(args, values...)
}

// implementation

type namedArgument struct {
	litsql.ArgumentBase
	name string
}

func (a namedArgument) Name() string {
	return a.name
}

type namedArgumentWithDefault struct {
	litsql.ArgumentBase
	name         string
	defaultValue any
}

func (a namedArgumentWithDefault) Name() string {
	return a.name
}

func (a namedArgumentWithDefault) Value() (any, error) {
	return a.defaultValue, nil
}

type dbNamedArgument struct {
	litsql.ArgumentBase
	name string
}

func (a dbNamedArgument) DBName() string {
	return a.name
}

type dbNamedArgumentWithDefault struct {
	litsql.ArgumentBase
	name         string
	defaultValue any
}

func (a dbNamedArgumentWithDefault) DBName() string {
	return a.name
}

func (a dbNamedArgumentWithDefault) Value() (any, error) {
	return a.defaultValue, nil
}

type funcArgument struct {
	litsql.ArgumentBase
	fn func() (any, error)
}

func (f funcArgument) Value() (any, error) {
	return f.fn()
}
