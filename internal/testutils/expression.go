package testutils

import (
	"bytes"
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal"
	"gotest.tools/v3/assert"
)

func TestExpression(t *testing.T, e litsql.Expression, output string, args ...any) {
	t.Helper()

	var buf bytes.Buffer
	w := internal.NewWriter(&buf,
		internal.WithWriterUseNewLine(false),
	)

	gotArgs, err := litsql.Express(w, NewTestDialect(), 1, e)
	assert.NilError(t, err)
	assert.DeepEqual(t, output, buf.String())
	assert.DeepEqual(t, args, gotArgs)
}

func TestExpressionSlice(t *testing.T, e []litsql.Expression, output string, args ...any) {
	t.Helper()

	var buf bytes.Buffer
	w := internal.NewWriter(&buf,
		internal.WithWriterUseNewLine(false),
	)

	gotArgs, err := litsql.ExpressSlice(w, NewTestDialect(), 1, e, nil, nil, nil)
	assert.NilError(t, err)
	assert.DeepEqual(t, output, buf.String())
	assert.DeepEqual(t, args, gotArgs)
}

func TestExpressionIsError(t *testing.T, e litsql.Expression) {
	t.Helper()
	TestExpressionErrorIs(t, e, nil)
}

func TestExpressionErrorIs(t *testing.T, e litsql.Expression, errIs error) {
	t.Helper()
	var buf bytes.Buffer
	w := internal.NewWriter(&buf,
		internal.WithWriterUseNewLine(false),
	)

	_, err := litsql.Express(w, NewTestDialect(), 1, e)
	if errIs != nil {
		assert.ErrorIs(t, err, errIs)
	} else {
		assert.Assert(t, err != nil)
	}
}

func TestExpressionSliceIsError(t *testing.T, e []litsql.Expression) {
	t.Helper()
	TestExpressionSliceErrorIs(t, e, nil)
}

func TestExpressionSliceErrorIs(t *testing.T, e []litsql.Expression, errIs error) {
	t.Helper()
	var buf bytes.Buffer
	w := internal.NewWriter(&buf,
		internal.WithWriterUseNewLine(false),
	)

	_, err := litsql.ExpressSlice(w, NewTestDialect(), 1, e, nil, nil, nil)
	if errIs != nil {
		assert.ErrorIs(t, err, errIs)
	} else {
		assert.Assert(t, err != nil)
	}
}
