package internal

import (
	"database/sql"
	"fmt"

	"github.com/rrgmc/litsql"
)

// GetArgValuesInstance gets the [litsql.ArgValues] instance from the passed parameter.
func GetArgValuesInstance(values any) (litsql.ArgValues, error) {
	switch xv := values.(type) {
	case litsql.ArgValues:
		return xv, nil
	case map[string]any:
		return litsql.MapArgValues(xv), nil
	default:
		return nil, fmt.Errorf("unsupported arg values type: %T", values)
	}
}

// ParseArgs replaces all [litsql.Argument] instances in args with named values.
func ParseArgs(args []any, values any) ([]any, error) {
	av, err := GetArgValuesInstance(values)
	if err != nil {
		return nil, err
	}
	return ParseArgValues(args, av)
}

// ParseArgValues replaces all [litsql.Argument] instances in args with named values.
func ParseArgValues(args []any, values litsql.ArgValues) ([]any, error) {
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

			if values == nil {
				nok = false
			} else {
				v, nok = values.Get(atname)
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
