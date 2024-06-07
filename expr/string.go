package expr

import (
	"github.com/rrgmc/litsql"
)

// String returns a raw string expression.
func String(str string) litsql.Expression {
	return Raw(str)
}

// StringQuote returns a quoted string expression.
func StringQuote(str string) litsql.Expression {
	return Quote(str)
}

// StringQuoteCheck returns a checked quoted string expression (only quote if needed).
func StringQuoteCheck(str string) litsql.Expression {
	return QuoteCheck(str)
}

// StringList converts a slice of strings to a slice of raw string expressions.
func StringList(str []string) []litsql.Expression {
	var ret []litsql.Expression
	for _, s := range str {
		ret = append(ret, String(s))
	}
	return ret
}

// StringListQuote converts a slice of strings to a slice of quoted raw string expressions.
func StringListQuote(str []string) []litsql.Expression {
	var ret []litsql.Expression
	for _, s := range str {
		ret = append(ret, Quote(s))
	}
	return ret
}

// StringListQuoteCheck converts a slice of strings to a slice of checked quoted raw string expressions.
func StringListQuoteCheck(str []string) []litsql.Expression {
	var ret []litsql.Expression
	for _, s := range str {
		ret = append(ret, QuoteCheck(s))
	}
	return ret
}
