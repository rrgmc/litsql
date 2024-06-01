package litsql

import (
	"database/sql"
	"fmt"
)

// Argument is the base interface for query arguments.
type Argument interface {
	isArgument()
}

// NamedArgument represents an argument were its value will be provided by name.
type NamedArgument interface {
	Argument
	Name() string
}

// ValuedArgument represents an argument were its value will be provided by this instance.
type ValuedArgument interface {
	Argument
	Value() (any, error)
}

// DBNamedArgument is like NamedArgument but its value will be wrapped using [sql.Named].
type DBNamedArgument interface {
	Argument
	DBName() string
}

type namedArgument struct {
	ArgumentBase
	name string
}

func (a namedArgument) Name() string {
	return a.name
}

type namedArgumentWithDefault struct {
	ArgumentBase
	name         string
	defaultValue any
}

func (a namedArgumentWithDefault) Name() string {
	return a.name
}

func (f namedArgumentWithDefault) Value() (any, error) {
	return f.defaultValue, nil
}

type dbNamedArgument struct {
	ArgumentBase
	name string
}

func (a dbNamedArgument) DBName() string {
	return a.name
}

type funcArgument struct {
	ArgumentBase
	fn func() (any, error)
}

func (f funcArgument) Value() (any, error) {
	return f.fn()
}

func Arg(name string) Argument {
	return namedArgument{name: name}
}

func ArgDefault(name string, defaultValue any) Argument {
	return namedArgumentWithDefault{name: name, defaultValue: defaultValue}
}

func DBArg(name string) DBNamedArgument {
	return dbNamedArgument{name: name}
}

func ArgFunc(fn func() (any, error)) Argument {
	return funcArgument{fn: fn}
}

type ArgValues interface {
	Get(string) (any, bool)
}

type MapArgValues map[string]any

func (m MapArgValues) Get(s string) (any, bool) {
	v, ok := m[s]
	return v, ok
}

func ParseArgs(args []any, values ...ArgValues) ([]any, error) {
	var ret []any
	for _, arg := range args {
		// if both Named and Valued, check Named first, Valued as a fallback.
		atvalued, isvalued := arg.(ValuedArgument)

		atname := ""
		isatname := false
		isatdbnamed := false

		if at, ok := arg.(NamedArgument); ok {
			atname = at.Name()
			isatname = true
		} else if at, ok := arg.(DBNamedArgument); ok {
			atname = at.DBName()
			isatname = true
			isatdbnamed = true
		}

		if isatname {
			var nok bool
			var v any

			if len(values) == 0 {
				nok = false
			} else {
				for _, value := range values {
					v, nok = value.Get(atname)
					if nok {
						break
					}
				}
			}
			if !nok {
				if !isvalued {
					return nil, fmt.Errorf("value for argument '%s' not found", atname)
				}
			} else {
				if isatdbnamed {
					v = sql.Named(atname, v)
				}
				ret = append(ret, v)
				continue
			}
		}

		if isvalued {
			v, err := atvalued.Value()
			if err != nil {
				return nil, err
			}
			if isatdbnamed {
				v = sql.Named(atname, v)
			}
			ret = append(ret, v)
			continue
		}

		ret = append(ret, arg)
	}
	return ret, nil
}

// helpers

type ArgumentBase struct{}

func (a ArgumentBase) isArgument() {}
