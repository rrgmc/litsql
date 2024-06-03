package internal

import (
	"fmt"

	"github.com/rrgmc/litsql"
)

// MergeClauses merge all clauses into the first one in the list.
func MergeClauses[T litsql.QueryClause](clauses ...T) (T, error) {
	var c T
	var cmerge litsql.QueryClauseMerge
	for ci, clause := range clauses {
		cm, ok := any(clause).(litsql.QueryClauseMerge)
		if !ok {
			return c, fmt.Errorf("clause %d is not QueryClauseMerge", ci)
		}

		if ci == 0 {
			c = clause
			cmerge = cm
		} else {
			err := cmerge.ClauseMerge(clause)
			if err != nil {
				return c, err
			}
		}
	}
	return c, nil
}
