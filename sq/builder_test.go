package sq

import (
	"bytes"
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/expr"
	"github.com/rrgmc/litsql/internal/testutils"
	"gotest.tools/v3/assert"
)

func TestBuilder(t *testing.T) {
	d := testutils.NewTestDialect()
	qb := NewQueryBuilder(d)
	qb.AddQueryClause(&testQueryClause{
		corder: 10,
		cid:    "c10",
		e:      expr.Raw("x"),
	})
	qb.AddQueryClause(&testQueryClause{
		corder: 5,
		cid:    "c5",
		e:      expr.Raw("y"),
	})
	builderTest(t, d, qb, "yx")
}

func TestBuilderArgs(t *testing.T) {
	d := testutils.NewTestDialect()
	qb := NewQueryBuilder(d)
	qb.AddQueryClause(&testQueryClause{
		corder: 10,
		cid:    "c10",
		e:      expr.Clause("x?", 12),
	})
	qb.AddQueryClause(&testQueryClause{
		corder: 5,
		cid:    "c5",
		e:      expr.Raw("y"),
	})
	builderTest(t, d, qb, "yx$1",
		12)
}

func TestBuilderMergeMultiple(t *testing.T) {
	qmergefunc := func(this, other litsql.QueryClause) error {
		xthis, thisok := this.(*testQueryClauseMerge)
		xother, otherok := other.(*testQueryClauseMerge)
		assert.Assert(t, thisok, "'this' query clause is not of the expected type")
		assert.Assert(t, otherok, "'other' query clause is not of the expected type")
		xthis.e = expr.Join(xthis.e, xother.e)
		return nil
	}

	d := testutils.NewTestDialect()
	qb := NewQueryBuilder(d)
	qb.AddQueryClause(&testQueryClauseMultiple{
		testQueryClause: testQueryClause{
			corder: 2,
			cid:    "c2",
			e:      expr.Raw("M2_"),
		},
	})
	qb.AddQueryClause(&testQueryClause{
		corder: 10,
		cid:    "c10",
		e:      expr.Raw("C10_"),
	})
	qb.AddQueryClause(&testQueryClauseMultiple{
		testQueryClause: testQueryClause{
			corder: 9,
			cid:    "c9",
			e:      expr.Raw("M9_"),
		},
	})
	qb.AddQueryClause(&testQueryClauseMerge{
		testQueryClause: testQueryClause{
			corder: 8,
			cid:    "c8",
			e:      expr.Raw("R8_"),
		},
		m: qmergefunc,
	})
	qb.AddQueryClause(&testQueryClauseMultiple{
		testQueryClause: testQueryClause{
			corder: 9,
			cid:    "c9",
			e:      expr.Raw("M9_"),
		},
	})
	qb.AddQueryClause(&testQueryClause{
		corder: 5,
		cid:    "c5",
		e:      expr.Raw("C5_"),
	})
	qb.AddQueryClause(&testQueryClauseMerge{
		testQueryClause: testQueryClause{
			corder: 8,
			cid:    "c8",
			e:      expr.Raw("R8_"),
		},
		m: qmergefunc,
	})

	builderTest(t, d, qb, "M2_C5_R8_R8_M9_M9_C10_")
}

func TestBuilderMergeNoMultiple(t *testing.T) {
	d := testutils.NewTestDialect()
	qb := NewQueryBuilder(d)
	qb.AddQueryClause(&testQueryClause{
		corder: 10,
		cid:    "c10",
		e:      expr.Raw("C10_"),
	})
	qb.AddQueryClause(&testQueryClause{
		corder: 10,
		cid:    "c10",
		e:      expr.Raw("C102_"),
	})

	_, err := qb.QueryClauseList()
	assert.ErrorIs(t, err, litsql.ErrClause)
}

func builderTest(t *testing.T, d litsql.Dialect, qb *Builder, querystr string, args ...any) {
	t.Helper()

	var buf bytes.Buffer
	w := NewWriter(&buf,
		WithUseNewLine(false))
	eb := litsql.NewExpressBuilder(w, d, 1)
	clauses, err := qb.QueryClauseList()
	assert.NilError(t, err)
	for _, e := range clauses {
		eb.Express(e)
	}

	assert.NilError(t, w.Err())
	gotArgs, err := eb.Result()
	assert.NilError(t, err)

	assert.Equal(t, querystr, buf.String())
	assert.DeepEqual(t, args, gotArgs)
}

type testQueryClause struct {
	cid    string
	corder int
	e      litsql.Expression
}

var _ litsql.QueryClause = (*testQueryClause)(nil)

func (c *testQueryClause) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) (args []any, err error) {
	return c.e.WriteSQL(w, d, start)
}

func (c *testQueryClause) ClauseID() string {
	return c.cid
}

func (c *testQueryClause) ClauseOrder() int {
	return c.corder
}

type testQueryClauseMerge struct {
	testQueryClause
	m func(this, other litsql.QueryClause) error
}

var _ litsql.QueryClauseMerge = (*testQueryClauseMerge)(nil)

func (c *testQueryClauseMerge) ClauseMerge(other litsql.QueryClause) error {
	return c.m(c, other)
}

type testQueryClauseMultiple struct {
	testQueryClause
}

var _ litsql.QueryClauseMultiple = (*testQueryClauseMultiple)(nil)

func (c *testQueryClauseMultiple) ClauseMultiple() {}
