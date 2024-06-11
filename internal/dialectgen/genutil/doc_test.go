package genutil

import (
	"go/ast"
	"testing"

	"gotest.tools/v3/assert"
)

func TestParseDoc(t *testing.T) {
	for _, test := range []struct {
		name       string
		doc        *ast.CommentGroup
		lines      []string
		directives Directives
	}{
		{
			name: "comment with space and directive",
			doc: &ast.CommentGroup{
				List: []*ast.Comment{
					{
						Text: "This is a comment",
					},
					{
						Text: "",
					},
					{
						Text: "litsql:dialects psql,sqlite",
					},
				},
			},
			lines: []string{
				"This is a comment",
			},
			directives: Directives{
				"dialects": Directive{
					Name:  "dialects",
					Value: "psql,sqlite",
				},
			},
		},
		{
			name: "comment with empty line",
			doc: &ast.CommentGroup{
				List: []*ast.Comment{
					{
						Text: "This is a comment",
					},
					{
						Text: "",
					},
				},
			},
			lines: []string{
				"This is a comment",
				"",
			},
			directives: Directives{},
		},
		{
			name: "only directive",
			doc: &ast.CommentGroup{
				List: []*ast.Comment{
					{
						Text: "litsql:dialects psql,sqlite",
					},
				},
			},
			lines: nil,
			directives: Directives{
				"dialects": Directive{
					Name:  "dialects",
					Value: "psql,sqlite",
				},
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			lines, directives := ParseDoc(test.doc)
			assert.DeepEqual(t, test.lines, lines)
			assert.DeepEqual(t, test.directives, directives)
		})
	}
}
