package internal_test

import (
	"testing"

	"github.com/rrgmc/litsql"
	"github.com/rrgmc/litsql/internal"
	"github.com/rrgmc/litsql/sq"
	"gotest.tools/v3/assert"
)

func TestBuildQuery(t *testing.T) {
	for _, test := range []struct {
		name          string
		f             func(w litsql.Writer, start int) (args []any, err error)
		expectedQuery string
		expectedArgs  []any
		options       []internal.BuildQueryOption
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
					&internal.NamedArgument{ArgName: "first"},
					&internal.NamedArgument{ArgName: "second"},
				}, nil
			},
			expectedQuery: "TEST QUERY WITH PARSED ARGUMENTS",
			expectedArgs:  []any{22, 44},
			options: []internal.BuildQueryOption{
				internal.WithBuildQueryParseArgs(sq.MapArgValues{
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
			options: []internal.BuildQueryOption{
				internal.WithBuildQueryWriterOptions(
					internal.WithWriterUseNewLine(true),
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
			options: []internal.BuildQueryOption{
				internal.WithBuildQueryWriterOptions(
					internal.WithWriterUseNewLine(false),
				),
			},
		},
		{
			name: "query with fixed ArgValues",
			f: func(w litsql.Writer, start int) (args []any, err error) {
				w.Write("TEST QUERY WITH FIXED ARGVALUES")
				return []any{
					&internal.NamedArgument{ArgName: "first"},
					&internal.NamedArgument{ArgName: "second"},
				}, nil
			},
			options: []internal.BuildQueryOption{
				internal.WithBuildQueryParseArgs(litsql.ArgValuesFunc(func(s string) (any, bool, error) {
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
			queryStr, args, err := internal.BuildQuery(query, test.options...)
			assert.NilError(t, err)
			assert.Equal(t, test.expectedQuery, queryStr)
			assert.DeepEqual(t, test.expectedArgs, args)
		})
	}
}
