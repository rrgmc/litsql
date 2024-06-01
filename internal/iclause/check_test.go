package iclause

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/sq"
	"gotest.tools/v3/assert"
)

func checkExpression(t *testing.T, e litsql.Expression, output testBuffer, args ...any) {
	t.Helper()

	for _, useNewLine := range []bool{false, true} {
		var buf bytes.Buffer
		w := sq.NewWriter(&buf,
			sq.WithWriterUseNewLine(useNewLine),
			sq.WithWriterIndentStr(" "),
		)

		gotArgs, err := e.WriteSQL(w, &testDialect{}, 1)
		assert.NilError(t, err)

		if useNewLine {
			assert.Equal(t, output.OutputNL(), buf.String())
		} else {
			assert.Equal(t, output.Output(), buf.String())
		}
		assert.DeepEqual(t, args, gotArgs)
	}
}

type testBuffer struct {
	b   bytes.Buffer
	bnl bytes.Buffer
}

func (b *testBuffer) Write(f string, args ...any) {
	_, _ = b.b.WriteString(fmt.Sprintf(f, args...))
	_, _ = b.bnl.WriteString(fmt.Sprintf(f, args...))
}

func (b *testBuffer) WriteNewLine() {
	_, _ = b.b.WriteString(" ")
	_, _ = b.bnl.WriteString("\n")
}

func (b *testBuffer) WriteIndent(amount int) {
	_, _ = b.b.WriteString("")
	_, _ = b.bnl.WriteString(strings.Repeat(" ", amount))
}

func (b *testBuffer) Output() string {
	return b.b.String()
}

func (b *testBuffer) OutputNL() string {
	return b.bnl.String()
}

type testDialect struct{}

func (d testDialect) WriteArg(w litsql.Writer, position int) {
	w.Write("$")
	w.Write(strconv.Itoa(position))
}

func (d testDialect) WriteQuoted(w litsql.Writer, s string) {
	w.Write(`"`)
	w.Write(s)
	w.Write(`"`)
}

func (d testDialect) WriteCheckQuoted(w litsql.Writer, s string) {
	if !strings.ContainsAny(s, " ") {
		w.Write(s)
		return
	}
	d.WriteQuoted(w, s)
}
