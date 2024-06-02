package testutils

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/rrgmc/litsql"
)

type TestBuffer struct {
	b   bytes.Buffer
	bnl bytes.Buffer
}

func NewTestBuffer() *TestBuffer {
	return &TestBuffer{}
}

func (b *TestBuffer) WriteTestClausePrefix() {
	_, _ = b.b.WriteString("@")
	_, _ = b.bnl.WriteString("@")
}

func (b *TestBuffer) Write(f string, args ...any) {
	_, _ = b.b.WriteString(fmt.Sprintf(f, args...))
	_, _ = b.bnl.WriteString(fmt.Sprintf(f, args...))
}

func (b *TestBuffer) WriteSeparator() {
	_, _ = b.b.WriteString(" ")
	_, _ = b.bnl.WriteString("\n")
}

func (b *TestBuffer) WriteNewLine() {
	_, _ = b.b.WriteString("")
	_, _ = b.bnl.WriteString("\n")
}

func (b *TestBuffer) WriteIndent(amount int) {
	_, _ = b.b.WriteString("")
	_, _ = b.bnl.WriteString(strings.Repeat(" ", amount))
}

func (b *TestBuffer) Output() string {
	return b.b.String()
}

func (b *TestBuffer) OutputNL() string {
	return b.bnl.String()
}

type TestDialect struct{}

func NewTestDialect() *TestDialect {
	return &TestDialect{}
}

func (d TestDialect) WriteArg(w litsql.Writer, position int) {
	w.Write("$")
	w.Write(strconv.Itoa(position))
}

func (d TestDialect) WriteQuoted(w litsql.Writer, s string) {
	w.Write(`"`)
	w.Write(s)
	w.Write(`"`)
}

func (d TestDialect) WriteCheckQuoted(w litsql.Writer, s string) {
	if !strings.ContainsAny(s, " ") {
		w.Write(s)
		return
	}
	d.WriteQuoted(w, s)
}
