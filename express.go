package litsql

func Express(w Writer, d Dialect, start int, e Expression) ([]any, error) {
	if e == nil {
		return nil, nil
	}
	return e.WriteSQL(w, d, start)
}

// ExpressIf expands an express if the condition evaluates to true
// it can also add a prefix and suffix
func ExpressIf(w Writer, d Dialect, start int, e Expression, cond bool, prefix, suffix Expression) ([]any, error) {
	if !cond {
		return nil, nil
	}
	h := NewExpressBuilder(w, d, start)
	h.Express(prefix)
	h.Express(e)
	h.Express(suffix)
	return h.Result()
}

// ExpressSlice is used to express a slice of expressions along with a prefix and suffix
func ExpressSlice(w Writer, d Dialect, start int, expressions []Expression, prefix, sep, suffix Expression) ([]any, error) {
	if len(expressions) == 0 {
		return nil, nil
	}
	h := NewExpressBuilder(w, d, start)
	h.Express(prefix)
	for k, e := range expressions {
		if k != 0 {
			h.Express(sep)
		}
		h.Express(e)
	}
	h.Express(suffix)
	return h.Result()
}
