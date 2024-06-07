package expr

import (
	"testing"

	"github.com/rrgmc/litsql/internal/testutils"
)

func TestIf(t *testing.T) {
	ex := If(true, Raw("test_me"))
	testutils.TestExpression(t, ex, "test_me")

	ex = If(false, Raw("test_me"))
	testutils.TestExpression(t, ex, "")
}

func TestIfElse(t *testing.T) {
	ex := IfElse(true, Raw("test1"), Raw("test2"))
	testutils.TestExpression(t, ex, "test1")

	ex = IfElse(false, Raw("test1"), Raw("test2"))
	testutils.TestExpression(t, ex, "test2")
}

func TestPrefixIf(t *testing.T) {
	ex := PrefixIf(true, Raw("@"), Raw("test_me"))
	testutils.TestExpression(t, ex, "@test_me")

	ex = PrefixIf(false, Raw("@"), Raw("test_me"))
	testutils.TestExpression(t, ex, "test_me")
}
