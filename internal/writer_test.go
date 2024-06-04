package internal

import (
	"bytes"
	"testing"

	"gotest.tools/v3/assert"
)

func TestNewWriter(t *testing.T) {
	var buf bytes.Buffer

	w := NewWriter(&buf, WithWriterUseNewLine(false))

	w.AddSeparator(true)
	w.Write("a")
	w.WriteSeparator()
	w.Write("b")
	w.WriteNewLine()
	w.Write("c")
	w.AddSeparator(false)
	w.Write("d")
	w.AddSeparator(true)
	w.Write("e")
	w.StartQuery()
	w.AddSeparator(false)
	w.Write("f")
	w.Indent()
	w.Write("g")
	w.Dedent()
	w.Write("h")

	assert.NilError(t, w.Err())
	assert.Equal(t, "a bc d efgh", buf.String())
}

func TestNewWriterNewLine(t *testing.T) {
	var buf bytes.Buffer

	w := NewWriter(&buf,
		WithWriterUseNewLine(true),
		WithWriterIndentString(" "))

	w.AddSeparator(true)
	w.Write("a")
	w.WriteSeparator()
	w.Write("b")
	w.WriteNewLine()
	w.Write("c")
	w.AddSeparator(false)
	w.Write("d")
	w.AddSeparator(true)
	w.Write("e")
	w.StartQuery()
	w.AddSeparator(false)
	w.Write("f")
	w.Indent()
	w.Write("g")
	w.Dedent()
	w.Write("h")

	assert.NilError(t, w.Err())
	assert.Equal(t, "a\nb\nc d\nefgh", buf.String())
}
