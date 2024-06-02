package sq

import (
	"database/sql"
	"fmt"

	"github.com/rrgmc/litsql"
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

// ArgFunc returns the argument value in a callback.
func ArgFunc(fn func() (any, error)) litsql.Argument {
	return funcArgument{fn: fn}
}

// ArgValues is the supplier of values for named arguments.
type ArgValues interface {
	Get(string) (any, bool)
}

// Args wraps parseable argument results.
type Args []any

func (a Args) Parse(values ...any) ([]any, error) {
	return ParseArgs(a, values...)
}

func (a Args) ParseValues(values ...ArgValues) ([]any, error) {
	return ParseArgValues(a, values...)
}

// ParseArgs replaces all [litsql.Argument] instances in args with named values.
func ParseArgs(args []any, values ...any) ([]any, error) {
	var av []ArgValues
	for _, v := range values {
		switch xv := v.(type) {
		case ArgValues:
			av = append(av, xv)
		case map[string]any:
			av = append(av, MapArgValues(xv))
		default:
			return nil, fmt.Errorf("unsupported arg values type: %T", v)
		}
	}
	return ParseArgValues(args, av...)
}

// ParseArgValues replaces all [litsql.Argument] instances in args with named values.
func ParseArgValues(args []any, values ...ArgValues) ([]any, error) {
	var ret []any
	for _, arg := range args {
		// if both Named and Valued, check Named first, Valued as a fallback.
		atvalued, isvalued := arg.(litsql.ValuedArgument)

		atname := ""
		isatname := false
		isatdbnamed := false

		if at, ok := arg.(litsql.NamedArgument); ok {
			atname = at.Name()
			isatname = true
		} else if at, ok := arg.(litsql.DBNamedArgument); ok {
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

// implementation

// MapArgValues is an ArgValues backed from a map[string]any.
type MapArgValues map[string]any

func (m MapArgValues) Get(s string) (any, bool) {
	v, ok := m[s]
	return v, ok
}

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

func (f namedArgumentWithDefault) Value() (any, error) {
	return f.defaultValue, nil
}

type dbNamedArgument struct {
	litsql.ArgumentBase
	name string
}

func (a dbNamedArgument) DBName() string {
	return a.name
}

type funcArgument struct {
	litsql.ArgumentBase
	fn func() (any, error)
}

func (f funcArgument) Value() (any, error) {
	return f.fn()
}
