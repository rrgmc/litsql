package testutils

import "github.com/rrgmc/litsql"

func Merge[T litsql.QueryClauseMerge](clauses ...T) (T, error) {
	var c T
	for ci, clause := range clauses {
		if ci == 0 {
			c = clause
		} else {
			err := c.ClauseMerge(clause)
			if err != nil {
				return c, err
			}
		}
	}
	return c, nil
}
