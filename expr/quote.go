package expr

import (
	"github.com/rrgmc/litsql"
)

// Quote outputs quoted and joined, something like "users"."id".
func Quote(aa ...string) litsql.Expression {
	ss := make([]string, 0, len(aa))
	for _, v := range aa {
		if v == "" {
			continue
		}
		ss = append(ss, v)
	}
	return quoted{data: ss}
}

// quoted and joined... something like "users"."id"
type quoted struct {
	data []string
}

func (q quoted) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	if len(q.data) == 0 {
		return nil, nil
	}

	// wrap in parentheses and join with comma
	k := 0 // not using the loop index to avoid empty strings
	for _, a := range q.data {
		if a == "" {
			continue
		}

		if k != 0 {
			w.Write(".")
		}
		k++

		d.WriteQuoted(w, a)
	}

	return nil, nil
}
