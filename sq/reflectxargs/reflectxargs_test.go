package reflectxargs

import (
	"testing"

	"github.com/rrgmc/litsql/dialect/psql"
	"github.com/rrgmc/litsql/dialect/psql/im"
	"github.com/rrgmc/litsql/sq"
	"gotest.tools/v3/assert"
)

type sampleUser struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func TestNew(t *testing.T) {
	user := sampleUser{
		Id:   123,
		Name: "John Doe",
	}

	pargs := New(user)
	assert.Assert(t, pargs != nil)

	query := psql.Insert(
		im.Into("users", "id", "name"),
		im.ValuesAN("id", "name"),
	)
	queryStr, args, err := query.Build(
		sq.WithWriterOptions(sq.WithUseNewLine(false)),
		sq.WithParseArgs(pargs),
	)

	assert.NilError(t, err)
	assert.Equal(t, "INSERT INTO users (id, name) VALUES ($1, $2)", queryStr)
	assert.DeepEqual(t, []any{123, "John Doe"}, args)
}

func TestNewGetter(t *testing.T) {
	user := sampleUser{
		Id:   123,
		Name: "John Doe",
	}

	pargs := New(user)
	assert.Assert(t, pargs != nil)

	query := psql.Insert(
		im.Into("users", "id", "name"),
		im.ValuesAN("id", "name"),
	)
	queryStr, args, err := query.Build(
		sq.WithWriterOptions(sq.WithUseNewLine(false)),
		sq.WithParseArgs(pargs, WithGetArgsValuesOption()),
	)

	assert.NilError(t, err)
	assert.Equal(t, "INSERT INTO users (id, name) VALUES ($1, $2)", queryStr)
	assert.DeepEqual(t, []any{123, "John Doe"}, args)
}