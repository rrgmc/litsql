package expr

import (
	"github.com/rrgmc/litsql"
)

func S(str string) litsql.Expression {
	return Raw(str)
}

func SQ(str string) litsql.Expression {
	return Quote(str)
}

func SQC(str string) litsql.Expression {
	return QuoteCheck(str)
}

func SL(str []string) []litsql.Expression {
	var ret []litsql.Expression
	for _, s := range str {
		ret = append(ret, S(s))
	}
	return ret
}

func SLQ(str []string) []litsql.Expression {
	var ret []litsql.Expression
	for _, s := range str {
		ret = append(ret, Quote(s))
	}
	return ret
}

func SLQC(str []string) []litsql.Expression {
	var ret []litsql.Expression
	for _, s := range str {
		ret = append(ret, QuoteCheck(s))
	}
	return ret
}
