package expr

import (
	"testing"

	"github.com/rrgmc/litsql/internal/testutils"
)

func TestArg(t *testing.T) {
	ex := Arg(85)
	testutils.TestExpression(t, ex, "$1", 85)
}

func TestArgs(t *testing.T) {
	ex := Args([]any{85, 45})
	testutils.TestExpressionSlice(t, ex, "$1$2", 85, 45)
}

func TestIn(t *testing.T) {
	ex := In([]any{15, 20, 25})
	testutils.TestExpression(t, ex, "$1, $2, $3", 15, 20, 25)
}

func TestInT(t *testing.T) {
	ex := InT([]int{15, 20, 25})
	testutils.TestExpression(t, ex, "$1, $2, $3", 15, 20, 25)
}

func TestInP(t *testing.T) {
	ex := InP([]any{15, 20, 25})
	testutils.TestExpression(t, ex, "($1, $2, $3)", 15, 20, 25)
}

func TestInPT(t *testing.T) {
	ex := InPT([]int{15, 20, 25})
	testutils.TestExpression(t, ex, "($1, $2, $3)", 15, 20, 25)
}
