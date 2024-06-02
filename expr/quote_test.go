package expr

import (
	"testing"

	"github.com/rrgmc/litsql/internal/testutils"
)

func TestQuote(t *testing.T) {
	ex := Quote("test_me1", "test_me2")
	testutils.TestExpression(t, ex, `"test_me1"."test_me2"`)
}

func TestQuoteCheck(t *testing.T) {
	ex := QuoteCheck("test_me1", "test me2")
	testutils.TestExpression(t, ex, `test_me1."test me2"`)
}
