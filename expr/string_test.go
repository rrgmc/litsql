package expr

import (
	"testing"

	"github.com/rrgmc/litsql/internal/testutils"
)

func TestS(t *testing.T) {
	ex := S("test_me")
	testutils.TestExpression(t, ex, "test_me")
}

func TestSQ(t *testing.T) {
	ex := SQ("test_me")
	testutils.TestExpression(t, ex, `"test_me"`)
}

func TestSL(t *testing.T) {
	ex := SL([]string{"test_me1", "test_me2"})
	testutils.TestExpressionSlice(t, ex, "test_me1test_me2")
}

func TestSLQ(t *testing.T) {
	ex := SLQ([]string{"test_me1", "test_me2"})
	testutils.TestExpressionSlice(t, ex, `"test_me1""test_me2"`)
}

func TestSLQC(t *testing.T) {
	ex := SLQC([]string{"test_me1", "test me2"})
	testutils.TestExpressionSlice(t, ex, `test_me1"test me2"`)
}
