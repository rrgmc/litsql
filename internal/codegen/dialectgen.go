package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() error {

	currentDir := getCurrentDir()
	ismDir := filepath.Clean(filepath.Join(currentDir, "..", "ism"))

	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, ismDir, nil, parser.ParseComments|parser.SkipObjectResolution)
	if err != nil {
		return err
	}

	for _, pkg := range pkgs {
		fmt.Printf("%s %s %s\n", strings.Repeat("=", 15), pkg.Name, strings.Repeat("=", 15))
		for fn, file := range pkg.Files {
			if strings.HasSuffix(fn, "_test.go") {
				continue
			}

			fmt.Printf("%s %s %s\n", strings.Repeat("-", 15), fn, strings.Repeat("-", 15))

			for _, d := range file.Decls {
				switch dt := d.(type) {
				case *ast.GenDecl:
					fmt.Printf("GenDecl\n")
				case *ast.FuncDecl:
					fmt.Printf("Func: %s\n", dt.Name.Name)
				default:
					fmt.Printf("Unknown type: %T\n", dt)
				}
			}
		}
	}

	return nil
}

func getCurrentDir() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("could not determine current directory")
	}
	return filepath.Dir(filename)
}
