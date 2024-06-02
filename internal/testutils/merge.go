package testutils

import "github.com/rrgmc/litsql"

func Merge[T litsql.QueryClauseMerge](clauses ...T) T {
	var c T
	for ci, clause := range clauses {
		if ci == 0 {
			c = clause
		} else {
			c.ClauseMerge(clause)
		}
	}
	return c
}
