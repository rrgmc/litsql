package expr

import (
	"testing"

	"github.com/rrgmc/litsql/internal/testutils"
)

func TestJoin(t *testing.T) {
	ex := Join(Raw("test_me"), Raw("test2"))
	testutils.TestExpression(t, ex, "test_metest2")
}

func TestJoinSep(t *testing.T) {
	ex := JoinSep("--", Raw("test_me"), Raw("test2"))
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
