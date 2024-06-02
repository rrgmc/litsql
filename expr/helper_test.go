package expr

import (
	"testing"

	"github.com/rrgmc/litsql/internal/testutils"
)

func TestExprIf(t *testing.T) {
	ex := ExprIf(true, Raw("test_me"))
	testutils.TestExpression(t, ex, "test_me")

	ex = ExprIf(false, Raw("test_me"))
	testutils.TestExpression(t, ex, "")
}

func TestExprIfElse(t *testing.T) {
	ex := ExprIfElse(true, Raw("test1"), Raw("test2"))
	testutils.TestExpression(t, ex, "test1")

	ex = ExprIfElse(false, Raw("test1"), Raw("test2"))
	testutils.TestExpression(t, ex, "test2")
}

func TestPrefixIf(t *testing.T) {
	ex := PrefixIf(true, Raw("@"), Raw("test_me"))
	testutils.TestExpression(t, ex, "@test_me")

	ex = PrefixIf(false, Raw("@"), Raw("test_me"))
	testutils.TestExpression(t, ex, "test_me")
}
