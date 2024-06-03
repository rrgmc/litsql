package internal

import (
	"database/sql"
	"fmt"

	"github.com/rrgmc/litsql"
)

// ParseArgs replaces all [litsql.Argument] instances in args with named values.
func ParseArgs(args []any, values ...any) ([]any, error) {
	var av []litsql.ArgValues
	for _, v := range values {
		switch xv := v.(type) {
		case litsql.ArgValues:
			av = append(av, xv)
		case map[string]any:
			av = append(av, litsql.MapArgValues(xv))
		default:
			return nil, fmt.Errorf("unsupported arg values type: %T", v)
		}
	}
	return ParseArgValues(args, av...)
}

// ParseArgValues replaces all [litsql.Argument] instances in args with named values.
func ParseArgValues(args []any, values ...litsql.ArgValues) ([]any, error) {
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
