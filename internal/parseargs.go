package internal

import (
	"database/sql"
	"fmt"

	"github.com/rrgmc/litsql"
)

// ParseArgs replaces all [litsql.Argument] instances in args with named values.
// Use [litsql.MapArgValues] to use a "map[string]any" as source.
func ParseArgs(args []any, values litsql.ArgValues) ([]any, error) {
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
			var err error

			if values == nil {
				nok = false
			} else {
				v, nok, err = values.Get(atname)
			}
			if err != nil {
				return nil, fmt.Errorf("%w: %w", litsql.ErrParseArgs, err)
			}
			if !nok {
				if !isvalued {
					return nil, fmt.Errorf("%w: value for argument '%s' not found", litsql.ErrParseArgs, atname)
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
