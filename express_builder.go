package litsql

// ExpressBuilder builds arguments in a sequence of Express calls.
type ExpressBuilder struct {
	w     Writer
	d     Dialect
	start int
	args  []any
	err   error
}

func NewExpressBuilder(w Writer, d Dialect, start int) *ExpressBuilder {
	return &ExpressBuilder{
		w:     w,
		d:     d,
		start: start,
	}
}

func (h *ExpressBuilder) Express(e Expression) {
	if h.err != nil || e == nil {
		return
	}
	var newArgs []any
	newArgs, h.err = Express(h.w, h.d, h.start+len(h.args), e)
	if h.err != nil {
		return
	}
	h.args = append(h.args, newArgs...)
}

func (h *ExpressBuilder) ExpressIf(e Expression, cond bool, prefix, suffix Expression) {
	if h.err != nil || e == nil {
		return
	}
	var newArgs []any
	newArgs, h.err = ExpressIf(h.w, h.d, h.start+len(h.args), e, cond, prefix, suffix)
	if h.err != nil {
		return
	}
	h.args = append(h.args, newArgs...)
}

func (h *ExpressBuilder) ExpressSlice(expressions []Expression, prefix, sep, suffix Expression) {
	if h.err != nil || len(expressions) == 0 {
		return
	}
	var newArgs []any
	newArgs, h.err = ExpressSlice(h.w, h.d, h.start+len(h.args), expressions, prefix, sep, suffix)
	if h.err != nil {
		return
	}
	h.args = append(h.args, newArgs...)
}

func (h *ExpressBuilder) WriteQuery(e Query) {
	if h.err != nil || e == nil {
		return
	}
	var newArgs []any
	newArgs, h.err = e.WriteQuery(h.w, h.start+len(h.args))
	if h.err != nil {
		return
	}
	h.args = append(h.args, newArgs...)
}

func (h *ExpressBuilder) Err() error {
	return h.err
}

func (h *ExpressBuilder) Result() ([]any, error) {
	return h.args, h.err
}
