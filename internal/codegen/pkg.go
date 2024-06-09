package main

import (
	"fmt"
	"go/types"
	"os"
	"path/filepath"

	jen2 "github.com/dave/jennifer/jen"
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

	f := jen2.NewFile("sm")
	f.PackageComment("// Code generated by \"litsql\"; DO NOT EDIT.")

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
			sig := funcTyp.Type().(*types.Signature)
			fmt.Printf("%s\n", types.ObjectString(funcTyp, qual))

			f.Func().Id(funcTyp.Name()).
				ParamsFunc(func(pgroup *jen2.Group) {
					for k := 0; k < sig.Params().Len(); k++ {
						sigParam := sig.Params().At(k)
						pgroup.Id(jen.ParamName(k, sigParam)).Add(jen.GetQualCode(sigParam.Type()))
					}
				}).
				ParamsFunc(func(rgroup *jen2.Group) {
					for k := 0; k < sig.Results().Len(); k++ {
						sigParam := sig.Results().At(k)
						rgroup.Id(sigParam.Name()).Add(jen.GetQualCode(sigParam.Type()))
					}
				}).
				Block()

			f.Line()

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

	err = f.Render(os.Stdout)
	if err != nil {
		return err
	}

	return nil
}

func runPkgTest() error {
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

	return nil
}
