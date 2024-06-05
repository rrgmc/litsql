package structargs

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/sq"
)

// WithGetArgsValuesOption adds ArgValues to be parsed by [sq.ParseArgs].
func WithGetArgsValuesOption() sq.GetArgValuesInstanceOption {
	return sq.WithGetArgValuesInstanceOptionCustom(func(values any) (litsql.ArgValues, error) {
		return ArgValues(values)
	})
}
