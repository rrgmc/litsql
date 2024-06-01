package litsql

type Writer interface {
	Write(s string)
	WriteNewLine()
	WriteSeparator()
	AddSeparator(topLevel bool)
	StartQuery()
	Indent()
	Dedent()
	Err() error
}
