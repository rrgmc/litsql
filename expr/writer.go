package expr

import "github.com/rrgmc/litsql"

// WriterNewLine calls [litsql.Writer.WriteNewline].
var WriterNewLine = litsql.ExpressionFunc(func(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	w.WriteNewLine()
	return nil, nil
})

// WriterSeparator calls [litsql.Writer.WriteSeparator].
var WriterSeparator = litsql.ExpressionFunc(func(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	w.WriteSeparator()
	return nil, nil
})

// WriterAddSeparator calls [litsql.Writer.AddSeparator].
func WriterAddSeparator(topLevel bool) litsql.Expression {
	return litsql.ExpressionFunc(func(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
		w.AddSeparator(topLevel)
		return nil, nil
	})
}
