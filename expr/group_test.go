package expr

import (
	"testing"

	"github.com/rrgmc/litsql/internal/testutils"
)

func TestParen(t *testing.T) {
	ex := Paren("test_me1", "test_me2")
	testutils.TestExpression(t, ex, "(test_me1, test_me2)")
}
