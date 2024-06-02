package testutils

import (
	"bytes"
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/sq"
	"gotest.tools/v3/assert"
)

func TestExpression(t *testing.T, e litsql.Expression, output *TestBuffer, args ...any) {
	t.Helper()

	if output.hasTestClausePrefix {
		e = TestClausePrefix(e)
	}

	for _, useNewLine := range []bool{false, true} {
		var buf bytes.Buffer
		w := sq.NewWriter(&buf,
			sq.WithWriterUseNewLine(useNewLine),
			sq.WithWriterIndentStr(" "),
		)

		gotArgs, err := e.WriteSQL(w, &TestDialect{}, 1)
		assert.NilError(t, err)

		if useNewLine {
			assert.DeepEqual(t, output.OutputNL(), buf.String())
		} else {
			assert.DeepEqual(t, output.Output(), buf.String())
		}
		assert.DeepEqual(t, args, gotArgs)
	}
}
