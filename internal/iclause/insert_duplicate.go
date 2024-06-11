package iclause

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal"
	"github.com/rrgmc/litsql/sq/clause"
)

type InsertDuplicateKey struct {
	Set Set
}

func (c *InsertDuplicateKey) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	if len(c.Set.Set) == 0 {
		return nil, nil
	}

	b := litsql.NewExpressBuilder(w, d, start)

	w.AddSeparator(true)
	w.Write("ON DUPLICATE KEY UPDATE")
	if !c.Set.Starter {
		c.Set.Starter = true
	}
	if !c.Set.SkipClause {
		c.Set.SkipClause = true
	}
	b.ExpressIf(&c.Set, len(c.Set.Set) > 0, nil, nil)

	return b.Result()
}

var _ litsql.QueryClauseMerge = (*InsertDuplicateKey)(nil)

func (c *InsertDuplicateKey) ClauseID() string {
	return "3a267e57-dc17-43b0-81db-6ae4f833bece"
}

func (c *InsertDuplicateKey) ClauseOrder() int {
	return clause.OrderInsertDuplicateKey
}

func (c *InsertDuplicateKey) ClauseMerge(other litsql.QueryClause) error {
	o, ok := other.(*InsertDuplicateKey)
	if !ok {
		return internal.NewClauseErrorInvalidMerge("InsertDuplicateKey")
	}
	if c.Set.Starter != o.Set.Starter {
		return internal.NewClauseErrorInvalidMergeHasChanges("Set")
	}
	c.Set.Set = append(c.Set.Set, o.Set.Set...)
	return nil
}

func (c *InsertDuplicateKey) SetSet(assignment litsql.Expression) {
	c.Set.Set = append(c.Set.Set, assignment)
}
