package expr

import (
	"errors"
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal/testutils"
)

func TestFunc(t *testing.T) {
	ex := Func(func() (litsql.Expression, error) {
		return JoinSep("*", Raw("a1"), Raw("a2")), nil
	})
	testutils.TestExpression(t, ex, "a1*a2")
}

func TestFuncError(t *testing.T) {
	err := errors.New("test error")

	ex := Func(func() (litsql.Expression, error) {
		return nil, err
	})
	testutils.TestExpressionErrorIs(t, ex, err)
}
