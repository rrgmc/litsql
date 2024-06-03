package structargs

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

func TestArgValues(t *testing.T) {
	user := sampleUser{
		Id:   123,
		Name: "John Doe",
	}

	pargs, err := ArgValues(user)
	assert.NilError(t, err)

	query := psql.Insert(
		im.Into("users", "id", "name"),
		im.ValuesAN("id", "name"),
	)
	queryStr, args, err := query.Build(
		sq.WithBuildWriterOptions(sq.WithWriterUseNewLine(false)),
		sq.WithBuildParseArgs(pargs),
	)

	assert.NilError(t, err)
	assert.Equal(t, "INSERT INTO users (id, name) VALUES ($1, $2)", queryStr)
	assert.DeepEqual(t, []any{123, "John Doe"}, args)
}
