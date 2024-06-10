package iclause

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal"
	"github.com/rrgmc/litsql/sq/clause"
)

type InsertConflictUpdate struct {
	Do     string // DO NOTHING | DO UPDATE
	Target InsertConflictTarget
	Set    Set
	Where  Where
}

func (c *InsertConflictUpdate) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	if c.Do == "" {
		return nil, internal.NewClauseError("'ON CONFLICT DO' cannot be empty")
	}

	b := litsql.NewExpressBuilder(w, d, start)

	w.AddSeparator(true)
	w.Write("ON CONFLICT")
	b.ExpressIf(&c.Target, true, nil, nil)

	w.Write(" DO ")
	w.Write(c.Do)
	if !c.Set.Starter {
		c.Set.Starter = true
	}
	b.ExpressIf(&c.Set, len(c.Set.Set) > 0, nil, nil)
	b.ExpressIf(&c.Where, len(c.Where.Conditions) > 0, nil, nil)

	return b.Result()
}

var _ litsql.QueryClause = (*InsertConflictUpdate)(nil)

func (c *InsertConflictUpdate) ClauseID() string {
	return "4ba79d92-d9f8-4806-b62d-7ba1c3974d1f"
}

func (c *InsertConflictUpdate) ClauseOrder() int {
	return clause.OrderInsertConflict
}

func (c *InsertConflictUpdate) SetWhere(condition litsql.Expression) {
	c.Target.Where = append(c.Target.Where, condition)
}

func (c *InsertConflictUpdate) SetDoNothing() {
	c.Do = "NOTHING"
}

func (c *InsertConflictUpdate) SetDoUpdate() {
	c.Do = "UPDATE"
}

type InsertConflictTarget struct {
	Constraint string
	Columns    []string
	Where      []litsql.Expression
}

func (c *InsertConflictTarget) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	b := litsql.NewExpressBuilder(w, d, start)

	if c.Constraint != "" {
		w.Write(" ON CONSTRAINT ")
		w.Write(c.Constraint)
		return b.Result()
	}

	b.ExpressSlice(expr.StringList(c.Columns), expr.Raw(" ("), expr.CommaSpace, expr.ClosePar)
	b.ExpressSlice(c.Where, expr.Raw(" WHERE "), expr.Raw(" AND "), nil)

	return b.Result()
}
