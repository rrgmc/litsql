package testutils

import (
	"bytes"
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal"
	"gotest.tools/v3/assert"
)

func TestWriterExpression(t *testing.T, e litsql.Expression, output *TestBuffer, args ...any) {
	t.Helper()

	if output.hasTestClausePrefix {
		e = testClausePrefix(e)
	}

	for _, useNewLine := range []bool{false, true} {
		var buf bytes.Buffer
		w := internal.NewWriter(&buf,
			internal.WithWriterUseNewLine(useNewLine),
			internal.WithWriterIndentStr(" "),
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

func TestWriterExpressionIsError(t *testing.T, e litsql.Expression) {
	t.Helper()
	TestWriterExpressionErrorIs(t, e, nil)
}

func TestWriterExpressionErrorIs(t *testing.T, e litsql.Expression, errIs error) {
	t.Helper()
	for _, useNewLine := range []bool{false, true} {
		var buf bytes.Buffer
		w := internal.NewWriter(&buf,
			internal.WithWriterUseNewLine(useNewLine),
			internal.WithWriterIndentStr(" "),
		)

		_, err := e.WriteSQL(w, &TestDialect{}, 1)
		if errIs != nil {
			assert.ErrorIs(t, err, errIs)
		} else {
			assert.Assert(t, err != nil)
		}
	}
}
