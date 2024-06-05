package expr

import (
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal/testutils"
)

func TestC(t *testing.T) {
	ex := C("test_me = ?", 98)
	testutils.TestExpression(t, ex, "test_me = $1", 98)
}

func TestClause(t *testing.T) {
	for _, test := range []struct {
		name          string
		expr          litsql.Expression
		expected      string
		expectedArgs  []any
		expectedError error
	}{
		{
			name: "plain",
			expr: clause{
				query: "SELECT a, b FROM alphabet",
			},
			expected: `SELECT a, b FROM alphabet`,
		},
		{
			name: "escaped args",
			expr: clause{
				query: `SELECT a, b FROM "alphabet\?" WHERE c = ? AND d <= ?`,
				args:  []any{1, 2},
			},
			expected:     `SELECT a, b FROM "alphabet?" WHERE c = $1 AND d <= $2`,
			expectedArgs: []any{1, 2},
		},
		{
			name: "mismatched args and placeholders",
			expr: clause{
				query: "SELECT a, b FROM alphabet WHERE c = ? AND d <= ?",
			},
			expected:      `SELECT a, b FROM alphabet WHERE c = $1 AND d <= $2`,
			expectedError: &clauseError{args: 0, placeholders: 2},
		},
		{
			name: "numbered args",
			expr: clause{
				query: "SELECT a, b FROM alphabet WHERE c = ? AND d <= ?",
				args:  []any{1, 2},
			},
			expected:     `SELECT a, b FROM alphabet WHERE c = $1 AND d <= $2`,
			expectedArgs: []any{1, 2},
		},
		{
			name: "expr args",
			expr: clause{
				query: "SELECT a, b FROM alphabet WHERE c IN (?) AND d <= ?",
				args:  []any{In([]any{5, 6, 7}), 2},
			},
			expected:     `SELECT a, b FROM alphabet WHERE c IN ($1, $2, $3) AND d <= $4`,
			expectedArgs: []any{5, 6, 7, 2},
		},
		{
			name: "expr args group",
			expr: clause{
				query: "SELECT a, b FROM alphabet WHERE c IN ? AND d <= ?",
				args:  []any{InP([]any{5, 6, 7}), 2},
			},
			expected:     `SELECT a, b FROM alphabet WHERE c IN ($1, $2, $3) AND d <= $4`,
			expectedArgs: []any{5, 6, 7, 2},
		},
		{
			name: "expr args quote",
			expr: clause{
				query: "SELECT a, b FROM alphabet WHERE c = ? AND d <= ?",
				args:  []any{Quote("AA"), 2},
			},
			expected:     `SELECT a, b FROM alphabet WHERE c = "AA" AND d <= $1`,
			expectedArgs: []any{2},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			if test.expectedError != nil {
				testutils.TestExpressionErrorIs(t, test.expr, test.expectedError)
			} else {
				testutils.TestExpression(t, test.expr, test.expected, test.expectedArgs...)
			}
		})
	}
}
