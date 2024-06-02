package expr

import (
	"testing"

	"github.com/rrgmc/litsql/internal/testutils"
)

func TestC(t *testing.T) {
	ex := C("test_me = ?", 98)
	testutils.TestExpression(t, ex, "test_me = $1", 98)
}
