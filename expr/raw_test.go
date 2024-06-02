package expr

import (
	"testing"

	"github.com/rrgmc/litsql/internal/testutils"
)

func TestRaw(t *testing.T) {
	ex := Raw("test_me")
	testutils.TestExpression(t, ex, "test_me")
}
