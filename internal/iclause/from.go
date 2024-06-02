package iclause

import (
	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal"
	"github.com/rrgmc/litsql/sq/clause"
)

type From struct {
	Table   litsql.Expression
	Starter bool
	Clause  string

	// Aliases
	Alias        string
	AliasColumns []string

	// Dialect specific modifiers
	Only           bool // Postgres
	Lateral        bool // Postgres & MySQL
	WithOrdinality bool // Postgres
}

func (c *From) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) ([]any, error) {
	b := litsql.NewExpressBuilder(w, d, start)

	if c.Starter {
		cl := c.Clause
		if cl == "" {
			cl = "FROM"
		}
		w.AddSeparator(true)
		w.Write(cl)
		w.Write(" ")
	}

	if c.Table == nil {
		return b.Result()
	}

	if c.Only {
		w.Write("ONLY ")
	}
	if c.Lateral {
		w.Write("LATERAL ")
	}
	b.Express(c.Table)
	if c.WithOrdinality {
		w.Write(" WITH ORDINALITY")
	}

	if c.Alias != "" {
		w.Write(" AS ")
		w.Write(c.Alias)
	}

	if len(c.AliasColumns) > 0 {
		w.Write(internal.OpenPar)
		for k, cAlias := range c.AliasColumns {
			if k != 0 {
				w.Write(", ")
			}
			w.Write(cAlias)
		}
		w.Write(internal.ClosePar)
	}

	return b.Result()
}

var _ litsql.QueryClause = (*From)(nil)

func (c *From) ClauseID() string {
	return "03837326-2900-4676-88de-b9aee378a869"
}

func (c *From) ClauseOrder() int {
	return clause.OrderFrom
}

func (f *From) SetAs(alias string, columns ...string) {
	f.Alias = alias
	f.AliasColumns = columns
}

func (f *From) SetOnly() {
	f.Only = true
}

func (f *From) SetLateral() {
	f.Lateral = true
}

func (f *From) SetWithOrdinality() {
	f.WithOrdinality = true
}
