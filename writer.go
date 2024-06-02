package litsql

// Writer is the interface used by expressions to output strings.
type Writer interface {
	// Write writes a string.
	Write(s string)
	// WriteNewLine writes a newline if in newline-mode, or nothing if not.
	WriteNewLine()
	// WriteSeparator writes a newline if in newline-mode, or a space if not.
	WriteSeparator()
	// AddSeparator schedules a WriteSeparator to be written on the next Write, except on the first Write call.
	// If toplevel is true, will try to write a newline if enabled, if false will add a space.
	AddSeparator(topLevel bool)
	// StartQuery signals the writer that a new query (or subquery) will start. It resets the "first Write" flag.
	StartQuery()
	// Indent increases indentation by 1 (only in newline-mode).
	Indent()
	// Dedent decreases indentation by 1 (only in newline-mode).
	Dedent()
	// Err returns any errors that were generated in the write process.
	Err() error
}
