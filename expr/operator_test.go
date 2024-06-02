package expr

import (
	"testing"

	"github.com/rrgmc/litsql/internal/testutils"
)

func TestJ(t *testing.T) {
	ex := J(Raw("test_me"), Raw("test2"))
	testutils.TestExpression(t, ex, "test_metest2")
}

func TestJS(t *testing.T) {
	ex := JS("--", Raw("test_me"), Raw("test2"))
	testutils.TestExpression(t, ex, "test_me--test2")
}

func TestOr(t *testing.T) {
	ex := Or("test_me", "test2")
	testutils.TestExpression(t, ex, "test_me OR test2")
}

func TestAnd(t *testing.T) {
	ex := And("test_me", "test2")
	testutils.TestExpression(t, ex, "test_me AND test2")
}
