package main

import (
	"cmp"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"slices"
	"strings"
)

func runAst() error {
	currentDir := getCurrentDir()
	ismDir := filepath.Clean(filepath.Join(currentDir, "..", "ism"))
	// dialectDir := filepath.Clean(filepath.Join(currentDir, "..", "..", "dialect"))
	// psqSmlDir := filepath.Join(dialectDir, "psql")

	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, ismDir, nil, parser.ParseComments|parser.SkipObjectResolution)
	if err != nil {
		return err
	}

	var funcs []*ast.FuncDecl

	for _, pkg := range pkgs {
		fmt.Printf("%s %s %s\n", strings.Repeat("=", 15), pkg.Name, strings.Repeat("=", 15))
		for fn, file := range pkg.Files {
			if strings.HasSuffix(fn, "_test.go") {
				continue
			}

			fmt.Printf("%s %s %s\n", strings.Repeat("-", 15), fn, strings.Repeat("-", 15))

			for _, d := range file.Decls {
				switch dt := d.(type) {
				// case *ast.GenDecl:
				// 	fmt.Printf("GenDecl\n")
				case *ast.FuncDecl:
					// fmt.Printf("Func: %s\n", dt.Name.Name)
					funcs = append(funcs, dt)
				default:
					fmt.Printf("Unknown type: %T\n", dt)
				}
			}
		}
	}

	slices.SortFunc(funcs, func(a, b *ast.FuncDecl) int {
		return cmp.Compare(a.Name.Name, b.Name.Name)
	})

	for _, funcDecl := range funcs {
		fmt.Printf("Func: %s\n", funcDecl.Name.Name)
	}

	return nil
}
