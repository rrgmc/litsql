package litsql

import "bytes"

func Build(q Query, writerOptions ...WriterOption) (string, []any, error) {
	var b bytes.Buffer
	w := NewWriter(&b, writerOptions...)
	args, err := q.WriteQuery(w, 1)
	if err != nil {
		return "", nil, err
	}
	if w.Err() != nil {
		return "", nil, err
	}
	return b.String(), args, nil
}
