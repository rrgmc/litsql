package sq

import (
	"io"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal"
)

// NewWriter creates the default implementation if [litsql.Writer].
func NewWriter(w io.Writer, options ...WriterOption) litsql.Writer {
	return internal.NewWriter(w, options...)
}

type WriterOption = internal.WriterOption

// WithWriterUseNewLine sets whether to use newlines in the output or not. Default is true.
func WithWriterUseNewLine(useNewLine bool) WriterOption {
	return internal.WithWriterUseNewLine(useNewLine)
}

// WithWriterIndentStr sets the indent string (used only if WithWriterUseNewLine is true). Default is "  " (two spaces).
func WithWriterIndentStr(indentStr string) WriterOption {
	return internal.WithWriterIndentStr(indentStr)
}
