package litsql

// Dialect implements dialect-specific methods.
type Dialect interface {
	WriteArg(w Writer, position int)
	WriteQuoted(w Writer, s string)
	WriteCheckQuoted(w Writer, s string) // quote only if string contains characters that need quoting.
}

// DialectWithNamed implements dialects that support db-specific named arguments.
type DialectWithNamed interface {
	Dialect
	WriteNamedArg(w Writer, name string)
}
