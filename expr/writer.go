package expr

import "github.com/rrgmc/litsql"

var WriterNewLine = litsql.ExpressionFunc(func(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	w.WriteNewLine()
	return nil, nil
})

var WriterSeparator = litsql.ExpressionFunc(func(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	w.WriteSeparator()
	return nil, nil
})

func WriterAddSeparator(topLevel bool) litsql.Expression {
	return litsql.ExpressionFunc(func(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
		w.AddSeparator(topLevel)
		return nil, nil
	})
}
