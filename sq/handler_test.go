package sq

import (
	"testing"

	"github.com/rrgmc/litsql"
	"gotest.tools/v3/assert"
)

func TestHandler(t *testing.T) {
	for _, test := range []struct {
		name          string
		f             func(w litsql.Writer, start int) (args []any, err error)
		expectedQuery string
		expectedArgs  []any
		options       []BuildOption
	}{
		{
			name: "query only",
			f: func(w litsql.Writer, start int) (args []any, err error) {
				w.Write("TEST QUERY")
				return nil, nil
			},
			expectedQuery: "TEST QUERY",
			expectedArgs:  nil,
		},
		{
			name: "query with arguments",
			f: func(w litsql.Writer, start int) (args []any, err error) {
				w.Write("TEST QUERY WITH ARGUMENTS")
				return []any{5, 8, 0}, nil
			},
			expectedQuery: "TEST QUERY WITH ARGUMENTS",
			expectedArgs:  []any{5, 8, 0},
		},
		{
			name: "query with parsed arguments",
			f: func(w litsql.Writer, start int) (args []any, err error) {
				w.Write("TEST QUERY WITH PARSED ARGUMENTS")
				return []any{
					NamedArg("first"),
					NamedArg("second"),
				}, nil
			},
			expectedQuery: "TEST QUERY WITH PARSED ARGUMENTS",
			expectedArgs:  []any{22, 44},
			options: []BuildOption{
				WithParseArgs(MapArgValues{
					"first":  22,
					"second": 44,
				}),
			},
		},
		{
			name: "query with writer use newline true",
			f: func(w litsql.Writer, start int) (args []any, err error) {
				w.Write("TEST QUERY WITH WRITER OPTIONS")
				w.WriteNewLine()
				return nil, nil
			},
			expectedQuery: "TEST QUERY WITH WRITER OPTIONS\n",
			options: []BuildOption{
				WithWriterOptions(
					WithUseNewLine(true),
				),
			},
		},
		{
			name: "query with writer use newline false",
			f: func(w litsql.Writer, start int) (args []any, err error) {
				w.Write("TEST QUERY WITH WRITER OPTIONS")
				w.WriteNewLine()
				return nil, nil
			},
			expectedQuery: "TEST QUERY WITH WRITER OPTIONS",
			options: []BuildOption{
				WithWriterOptions(
					WithUseNewLine(false),
				),
			},
		},
		{
			name: "query with fixed ArgValues",
			f: func(w litsql.Writer, start int) (args []any, err error) {
				w.Write("TEST QUERY WITH FIXED ARGVALUES")
				return []any{
					NamedArg("first"),
					NamedArg("second"),
				}, nil
			},
			options: []BuildOption{
				WithParseArgs(litsql.ArgValuesFunc(func(s string) (any, bool, error) {
					switch s {
					case "first":
						return 55, true, nil
					case "second":
						return 66, true, nil
					default:
						return nil, false, nil
					}
				})),
			},
			expectedQuery: "TEST QUERY WITH FIXED ARGVALUES",
			expectedArgs:  []any{55, 66},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			query := litsql.QueryFunc(nil, nil, test.f)
			h := NewHandler(WithDefaultBuildOptions(test.options...))
			queryStr, args, err := h.Build(query)
			assert.NilError(t, err)
			assert.Equal(t, test.expectedQuery, queryStr)
			assert.DeepEqual(t, test.expectedArgs, args)
		})
	}
}
