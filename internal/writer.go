package internal

import "io"

type WriterIO struct {
	w   io.Writer
	err error
}

func WewWriterIO(w io.Writer) *WriterIO {
	return &WriterIO{w: w}
}

func (w *WriterIO) Write(s string) {
	if w.err != nil {
		return
	}
	if len(s) == 0 {
		return
	}
	_, w.err = w.w.Write([]byte(s))
}

func (w *WriterIO) Err() error {
	return w.err
}
