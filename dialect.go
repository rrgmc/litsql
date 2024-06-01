package litsql

type Dialect interface {
	WriteArg(w Writer, position int)
	WriteQuoted(w Writer, s string)
	WriteCheckQuoted(w Writer, s string)
}

type DialectWithNamed interface {
	Dialect
	WriteNamedArg(w Writer, name string)
}
