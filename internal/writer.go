package internal

import (
	"io"
	"strings"

	"github.com/rrgmc/litsql"
)

// writer is the default implementation if [litsql.Writer].
type writer struct {
	wio              *WriterIO
	writeSep         bool
	writeSepTopLevel bool
	lastWrite        bool
	indent           int
	isNewLine        bool

	useNewLine bool
	indentStr  string
}

func NewWriter(w io.Writer, options ...WriterOption) litsql.Writer {
	ret := &writer{
		wio:        WewWriterIO(w),
		useNewLine: true,
		indentStr:  "  ",
	}
	for _, opt := range options {
		opt(ret)
	}
	return ret
}

type WriterOption func(*writer)

// WithWriterUseNewLine sets whether to use newlines in the output or not. Default is true.
func WithWriterUseNewLine(useNewLine bool) WriterOption {
	return func(w *writer) {
		w.useNewLine = useNewLine
	}
}

// WithWriterIndentString sets the indent string (used only if WithWriterUseNewLine is true). Default is "  " (two spaces).
func WithWriterIndentString(indentString string) WriterOption {
	return func(w *writer) {
		w.indentStr = indentString
	}
}

func (w *writer) Write(s string) {
	if w.wio.Err() != nil {
		return
	}
	if len(s) == 0 {
		return
	}
	if w.writeSep {
		if w.lastWrite {
			if w.writeSepTopLevel {
				w.WriteSeparator()
			} else {
				w.wio.Write(Space)
			}
		}
		w.writeSep = false
		w.writeSepTopLevel = false
		w.lastWrite = false
	}
	if w.isNewLine && w.indent > 0 {
		w.wio.Write(strings.Repeat(w.indentStr, w.indent))
	}
	w.lastWrite = true
	w.isNewLine = false
	w.wio.Write(s)
}

func (w *writer) WriteNewLine() {
	if w.useNewLine {
		w.wio.Write(NewLine)
		w.isNewLine = true
	}
}

func (w *writer) WriteSeparator() {
	if w.useNewLine {
		w.wio.Write(NewLine)
		w.isNewLine = true
	} else {
		w.wio.Write(Space)
	}
}

func (w *writer) AddSeparator(topLevel bool) {
	w.writeSep = true
	w.writeSepTopLevel = topLevel
}

func (w *writer) StartQuery() {
	w.lastWrite = false
}

func (w *writer) Indent() {
	w.indent += 1
}

func (w *writer) Dedent() {
	w.indent -= 1
	if w.indent < 0 {
		w.indent = 0
	}
}

func (w *writer) Err() error {
	return w.wio.Err()
}
