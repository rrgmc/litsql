package main

import (
	"fmt"
	"go/types"
	"path/filepath"

	"github.com/rrgmc/litsql/internal/codegen/jen"
	"golang.org/x/tools/go/packages"
)

func runPkg() error {
	currentDir := getCurrentDir()
	ismDir := filepath.Clean(filepath.Join(currentDir, "..", "ism"))
	// dialectDir := filepath.Clean(filepath.Join(currentDir, "..", "..", "dialect"))
	// psqSmlDir := filepath.Join(dialectDir, "psql")

	lpkg, err := jen.PkgInfoFromPath(
		ismDir, packages.NeedName|packages.NeedSyntax|packages.NeedTypes,
	)
	if err != nil {
		return fmt.Errorf("couldn't load source package: %s", err)
	}

	if lpkg.Types != nil {
		qual := types.RelativeTo(lpkg.Types)
		scope := lpkg.Types.Scope()
		for _, name := range scope.Names() {
			obj := scope.Lookup(name)
			if !obj.Exported() {
				continue // skip unexported names
			}
			funcTyp, ok := obj.(*types.Func)
			if !ok {
				continue
			}
			fmt.Printf("%s\n", types.ObjectString(funcTyp, qual))

			// fmt.Printf("%s\n", funcObj.FullName())
			// fmt.Printf("\t%s\n", types.ObjectString(obj, qual))
			// if tn, ok := obj.(*types.TypeName); ok {
			// 	fmt.Printf("%s (%s)\n", tn.Name(), types.ObjectString(obj, qual))
			//
			// 	for _, meth := range typeutil.IntuitiveMethodSet(obj.Type(), nil) {
			// 		if !meth.Obj().Exported() {
			// 			continue // skip unexported names
			// 		}
			// 		fmt.Printf("\t%s\n", types.SelectionString(meth, qual))
			// 	}
			// }
		}
	}

	// for _, objName := range srcPkg.Types.Scope().Names() {
	// 	fmt.Printf("%s\n", objName)
	// 	obj := srcPkg.Types.Scope().Lookup(objName)
	// 	if obj == nil {
	// 		fmt.Printf("%s object not found\n", objName)
	// 		continue
	// 	}
	// 	if ofunc, ok := obj.Type().(*types.Func); ok {
	//
	// 	}
	// }

	return nil
}
