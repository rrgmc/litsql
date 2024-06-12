package internal

import (
	"errors"
	"testing"

	"github.com/rrgmc/litsql"
	"gotest.tools/v3/assert"
)

func TestMergeClauses(t *testing.T) {
	c, err := MergeClauses(
		&testClauseMerge{id: 76},
		&testClauseMerge{id: 88},
		&testClauseMerge{id: 106},
	)
	assert.NilError(t, err)
	assert.Equal(t, 2, c.mergeCount)
	assert.Equal(t, 76, c.id) // merge is always done on the first item
}

func TestMergeClausesEmpty(t *testing.T) {
	c, err := MergeClauses[*testClauseMerge]()
	assert.NilError(t, err)
	assert.Assert(t, c == nil)
}

func TestMergeClausesError(t *testing.T) {
	me := errors.New("merge error")

	_, err := MergeClauses(
		&testClauseMerge{id: 76, mergeErr: me},
		&testClauseMerge{id: 88},
		&testClauseMerge{id: 106},
	)
	assert.ErrorIs(t, err, me)
}

func TestMergeClausesErrorOnlyInFirst(t *testing.T) {
	me := errors.New("merge error")

	_, err := MergeClauses(
		&testClauseMerge{id: 76},
		&testClauseMerge{id: 88, mergeErr: me}, // the ClauseMerge is only called on the first one
		&testClauseMerge{id: 106},
	)
	assert.NilError(t, err)
}

func TestMergeClausesNoMerge(t *testing.T) {
	_, err := MergeClauses[litsql.QueryClause](
		&testClauseMerge{id: 76},
		&testClauseNoMerge{},
		&testClauseMerge{id: 106},
	)
	assert.Assert(t, err != nil)
}

func TestMergeClausesNoMergeFirst(t *testing.T) {
	_, err := MergeClauses[litsql.QueryClause](
		&testClauseNoMerge{},
		&testClauseMerge{id: 76},
		&testClauseMerge{id: 106},
	)
	assert.Assert(t, err != nil)
}

type testClauseMerge struct {
	id         int
	mergeCount int
	mergeErr   error
}

func (t *testClauseMerge) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) (args []any, err error) {
	return nil, nil
}

func (t *testClauseMerge) ClauseID() string {
	return "testID"
}

func (t *testClauseMerge) ClauseOrder() int {
	return 51
}

func (t *testClauseMerge) ClauseMerge(other litsql.QueryClause) error {
	if t.mergeErr != nil {
		return t.mergeErr
	}
	t.mergeCount++
	return nil
}

type testClauseNoMerge struct {
}

func (t *testClauseNoMerge) WriteSQL(w litsql.Writer, d litsql.Dialect, start int) (args []any, err error) {
	return nil, nil
}

func (t *testClauseNoMerge) ClauseID() string {
	return "testID"
}

func (t *testClauseNoMerge) ClauseOrder() int {
	return 88
}
