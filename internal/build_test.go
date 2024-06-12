package internal

import (
	"testing"

	"github.com/rrgmc/litsql"
	"gotest.tools/v3/assert"
)

func TestBuildQuery(t *testing.T) {
	for _, test := range []struct {
		name          string
		f             func(w litsql.Writer, start int) (args []any, err error)
		expectedQuery string
		expectedArgs  []any
		options       []BuildQueryOption
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
					&NamedArgument{ArgName: "first"},
					&NamedArgument{ArgName: "second"},
				}, nil
			},
			expectedQuery: "TEST QUERY WITH PARSED ARGUMENTS",
			expectedArgs:  []any{22, 44},
			options: []BuildQueryOption{
				WithBuildQueryParseArgs(map[string]any{
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
			options: []BuildQueryOption{
				WithBuildQueryWriterOptions(
					WithWriterUseNewLine(true),
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
			options: []BuildQueryOption{
				WithBuildQueryWriterOptions(
					WithWriterUseNewLine(false),
				),
			},
		},
		{
			name: "query with custom ArgValues",
			f: func(w litsql.Writer, start int) (args []any, err error) {
				w.Write("TEST QUERY WITH CUSTOM ARGVALUES")
				return []any{
					&NamedArgument{ArgName: "first"},
					&NamedArgument{ArgName: "second"},
				}, nil
			},
			options: []BuildQueryOption{
				WithBuildQueryGetArgValuesInstanceOptions(
					WithGetArgValuesInstanceOptionCustom(func(values any) (litsql.ArgValues, error) {
						// for this test we are ignoring the input values
						return litsql.ArgValuesFunc(func(s string) (any, bool) {
							switch s {
							case "first":
								return 89, true
							case "second":
								return 61, true
							default:
								return nil, false
							}
						}), nil
					}),
				),
				// for this test we are ignoring the input values
				WithBuildQueryParseArgs(nil),
			},
			expectedQuery: "TEST QUERY WITH CUSTOM ARGVALUES",
			expectedArgs:  []any{89, 61},
		},
		{
			name: "query with fixed ArgValues",
			f: func(w litsql.Writer, start int) (args []any, err error) {
				w.Write("TEST QUERY WITH FIXED ARGVALUES")
				return []any{
					&NamedArgument{ArgName: "first"},
					&NamedArgument{ArgName: "second"},
				}, nil
			},
			options: []BuildQueryOption{
				WithBuildQueryParseArgValues(litsql.ArgValuesFunc(func(s string) (any, bool) {
					switch s {
					case "first":
						return 55, true
					case "second":
						return 66, true
					default:
						return nil, false
					}
				})),
			},
			expectedQuery: "TEST QUERY WITH FIXED ARGVALUES",
			expectedArgs:  []any{55, 66},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			query := litsql.QueryFunc(nil, nil, test.f)
			queryStr, args, err := BuildQuery(query, test.options...)
			assert.NilError(t, err)
			assert.Equal(t, test.expectedQuery, queryStr)
			assert.DeepEqual(t, test.expectedArgs, args)
		})
	}
}
