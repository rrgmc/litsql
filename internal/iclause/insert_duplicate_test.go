package iclause

import (
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/testutils"
)

func TestInsertDuplicateKey(t *testing.T) {
	clause := &InsertDuplicateKey{
		Set: Set{
			Set: []litsql.Expression{
				expr.Raw("id = 1"),
				expr.Raw("age = 50"),
			},
		},
	}

	o := testutils.NewTestBuffer()
	o.WriteSeparator()
	o.Write("ON DUPLICATE KEY UPDATE")

	o.WriteSeparator()
	o.Write("SET")
	o.WriteSeparator()
	o.WriteIndent(1)
	o.Write("id = 1,")
	o.WriteSeparator()
	o.WriteIndent(1)
	o.Write("age = 50")

	testutils.TestWriterExpression(t, clause, o)
}

func TestInsertDuplicateKeyEmpty(t *testing.T) {
	clause := &InsertDuplicateKey{}
	o := testutils.NewTestBuffer()
	testutils.TestWriterExpression(t, clause, o)
}
