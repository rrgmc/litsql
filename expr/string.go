package expr

import (
	"github.com/rrgmc/litsql"
)

// S returns a raw string expression.
func S(str string) litsql.Expression {
	return Raw(str)
}

// SQ returns a quoted string expression.
func SQ(str string) litsql.Expression {
	return Quote(str)
}

// SQC returns a checked quoted string expression (only quote if needed).
func SQC(str string) litsql.Expression {
	return QuoteCheck(str)
}

// SL converts a slice of strings to a slice of raw string expressions.
func SL(str []string) []litsql.Expression {
	var ret []litsql.Expression
	for _, s := range str {
		ret = append(ret, S(s))
	}
	return ret
}

// SLQ converts a slice of strings to a slice of quoted raw string expressions.
func SLQ(str []string) []litsql.Expression {
	var ret []litsql.Expression
	for _, s := range str {
		ret = append(ret, Quote(s))
	}
	return ret
}

// SLQC converts a slice of strings to a slice of checked quoted raw string expressions.
func SLQC(str []string) []litsql.Expression {
	var ret []litsql.Expression
	for _, s := range str {
		ret = append(ret, QuoteCheck(s))
	}
	return ret
}
