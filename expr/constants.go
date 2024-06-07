package expr

import (
	"errors"

	"github.com/rrgmc/litsql/internal"
)

var ErrNoNamedArgs = errors.New("Dialect does not support named arguments")

var (
	Space      = Raw(internal.Space)
	CommaSpace = Raw(internal.CommaSpace)
	NewLine    = Raw(internal.NewLine)
	OpenPar    = Raw(internal.OpenPar)
	ClosePar   = Raw(internal.ClosePar)

	CommaWriterNewLine   = Join(Raw(internal.Comma), WriterNewLine)
	CommaWriterSeparator = Join(Raw(internal.Comma), WriterSeparator)
)

//nolint:gochecknoglobals
var (
	null = Raw("NULL")
)
