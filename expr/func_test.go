package expr

import (
	"errors"
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal/testutils"
)

func TestF(t *testing.T) {
	ex := F(func() (litsql.Expression, error) {
		return JS("*", Raw("a1"), Raw("a2")), nil
	})
	testutils.TestExpression(t, ex, "a1*a2")
}

func TestFError(t *testing.T) {
	err := errors.New("test error")

	ex := F(func() (litsql.Expression, error) {
		return nil, err
	})
	testutils.TestExpressionErrorIs(t, ex, err)
}
