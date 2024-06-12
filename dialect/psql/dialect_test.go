package psql

import (
	"bytes"
	"testing"

	"github.com/rrgmc/litsql/sq"
	"gotest.tools/v3/assert"
)

func TestDialectWriteArg(t *testing.T) {
	var buf bytes.Buffer
	w := sq.NewWriter(&buf)
	Dialect.WriteArg(w, 3)

	assert.Equal(t, `$3`, buf.String())
}

func TestDialectWriteQuoted(t *testing.T) {
	var buf bytes.Buffer
	w := sq.NewWriter(&buf)
	Dialect.WriteQuoted(w, "a")

	assert.Equal(t, `"a"`, buf.String())
}
