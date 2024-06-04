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

// WithUseNewLine sets whether to use newlines in the output or not. Default is true.
func WithUseNewLine(useNewLine bool) WriterOption {
	return internal.WithWriterUseNewLine(useNewLine)
}

// WithIndentString sets the indent string (used only if WithUseNewLine is true). Default is "  " (two spaces).
func WithIndentString(indentString string) WriterOption {
	return internal.WithWriterIndentString(indentString)
}
