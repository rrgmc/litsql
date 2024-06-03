package testutils

import (
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal"
	"gotest.tools/v3/assert"
)

func TestQuery(t *testing.T, query litsql.Query, expected string, expectedArgs ...any) {
	t.Helper()
	queryStr, args, err := internal.BuildQuery(query, internal.WithBuildQueryWriterOptions(internal.WithWriterUseNewLine(false)))
	assert.NilError(t, err)
	assert.Equal(t, expected, queryStr)
	assert.DeepEqual(t, expectedArgs, args)
}

func TestQueryParseArgs(t *testing.T, query litsql.Query, expected string, argValues map[string]any, expectedArgs ...any) {
	t.Helper()
	queryStr, args, err := internal.BuildQuery(query,
		internal.WithBuildQueryWriterOptions(internal.WithWriterUseNewLine(false)),
		internal.WithBuildQueryParseArgs(argValues),
	)
	assert.NilError(t, err)
	assert.Equal(t, expected, queryStr)
	assert.DeepEqual(t, expectedArgs, args)
}

func TestQueryIsError(t *testing.T, query litsql.Query) {
	t.Helper()
	_, _, err := internal.BuildQuery(query, internal.WithBuildQueryWriterOptions(internal.WithWriterUseNewLine(false)))
	assert.Assert(t, err != nil)
}

func TestQueryErrorIs(t *testing.T, query litsql.Query, errIs error) {
	t.Helper()
	_, _, err := internal.BuildQuery(query, internal.WithBuildQueryWriterOptions(internal.WithWriterUseNewLine(false)))
	assert.ErrorIs(t, err, errIs)
}
