package sq

import (
	"database/sql"
	"fmt"

	"github.com/rrgmc/litsql"
)

func Arg(name string) litsql.Argument {
	return namedArgument{name: name}
}

func ArgDefault(name string, defaultValue any) litsql.Argument {
	return namedArgumentWithDefault{name: name, defaultValue: defaultValue}
}

func DBArg(name string) litsql.DBNamedArgument {
	return dbNamedArgument{name: name}
}

func ArgFunc(fn func() (any, error)) litsql.Argument {
	return funcArgument{fn: fn}
}

type ArgValues interface {
	Get(string) (any, bool)
}

func ParseArgs(args []any, values ...ArgValues) ([]any, error) {
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
