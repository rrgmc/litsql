package expr

import (
	"testing"

	"github.com/rrgmc/litsql/internal/testutils"
)

func TestS(t *testing.T) {
	ex := String("test_me")
	testutils.TestExpression(t, ex, "test_me")
}

func TestSQ(t *testing.T) {
	ex := StringQuote("test_me")
	testutils.TestExpression(t, ex, `"test_me"`)
}

func TestSL(t *testing.T) {
	ex := StringList([]string{"test_me1", "test_me2"})
	testutils.TestExpressionSlice(t, ex, "test_me1test_me2")
}

func TestSLQ(t *testing.T) {
	ex := StringListQuote([]string{"test_me1", "test_me2"})
	testutils.TestExpressionSlice(t, ex, `"test_me1""test_me2"`)
}

func TestSLQC(t *testing.T) {
	ex := StringListQuoteCheck([]string{"test_me1", "test me2"})
	testutils.TestExpressionSlice(t, ex, `test_me1"test me2"`)
}
