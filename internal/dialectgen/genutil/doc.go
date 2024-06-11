package genutil

import (
	"go/ast"
	"regexp"
	"slices"
	"strings"
)

type Directives map[string]Directive

func (d Directives) IsListValue(directive string, value string) bool {
	di, ok := d[directive]
	if !ok {
		return false
	}
	return di.IsListValue(value)
}

type Directive struct {
	Name  string
	Value string
}

func (d Directive) IsListValue(value string) bool {
	return slices.Contains(strings.Split(d.Value, ","), value)
}

func ParseDoc(doc *ast.CommentGroup) ([]string, Directives) {
	var lines []string
	directives := Directives{}

	for _, docLine := range doc.List {
		dmatches := reDirective.FindStringSubmatch(docLine.Text)
		if dmatches != nil {
			if len(dmatches) > 3 {
				directives[dmatches[2]] = Directive{
					Name:  dmatches[2],
					Value: dmatches[3],
				}
			}
		} else {
			lines = append(lines, docLine.Text)
		}
	}

	// strip trailing empty lines if there is directives
	lastline := len(lines) - 1
	if len(directives) > 0 {
		for ; lastline >= 0; lastline-- {
			if strings.TrimSpace(lines[lastline]) != "" {
				break
			}
		}
	}

	return lines[0 : lastline+1], directives
}

var (
	reDirective = regexp.MustCompile(`^([a-z0-9]+):([a-z0-9]+) (.+)$`)
)
