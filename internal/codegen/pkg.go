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
	sdir := "sm"
	sname := "Select"
	sdialect := "psql"

	currentDir := getCurrentDir()
	ismDir := filepath.Clean(filepath.Join(currentDir, "..", "i"+sdir))
	// dialectDir := filepath.Clean(filepath.Join(currentDir, "..", "..", "dialect"))
	// psqSmlDir := filepath.Join(dialectDir, "psql")

	ispkg := "github.com/rrgmc/litsql/internal/i" + sdir
	isqpkg := "github.com/rrgmc/litsql/internal/isq"
	sqpkg := "github.com/rrgmc/litsql/sq"
	sqchainpkg := "github.com/rrgmc/litsql/sq/chain"
	sdialectpkg := fmt.Sprintf("github.com/rrgmc/litsql/dialect/%s", sdialect)
	sdialecttagpkg := fmt.Sprintf("%s/tag", sdialectpkg)

	lpkg, err := jen.PkgInfoFromPath(
		ismDir, packages.NeedName|packages.NeedSyntax|packages.NeedTypes,
	)
	if err != nil {
		return fmt.Errorf("couldn't load source package: %s", err)
	}

	f := jen2.NewFile(sdir)
	f.PackageComment("// Code generated by \"litsql\"; DO NOT EDIT.")

	if lpkg.Types != nil {
		customNamedType := func(st jen2.Statement, tt *types.Named) *jen2.Statement {
			if tt.Obj().Name() == "QueryMod" && tt.Obj().Pkg().Path() == sqpkg {
				return st.Add(jen2.Qual(sdialectpkg, sname+"Mod"))
			} else if tt.Obj().Name() == "QueryModApply" && tt.Obj().Pkg().Path() == sqpkg {
				return st.Add(jen2.Qual(sdialectpkg, sname+"ModApply"))
			} else if tt.Obj().Name() == "Query" && tt.Obj().Pkg().Path() == isqpkg {
				return st.Add(jen2.Qual(sdialectpkg, sname+"Query"))
			} else if tt.Obj().Pkg().Name() == "chain" && tt.Obj().Pkg().Path() == sqchainpkg {
				return st.Add(jen2.Id(fmt.Sprintf("%sChain", tt.Obj().Name())))
			}
			return nil
		}

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
			// fmt.Printf("%s\n", types.ObjectString(funcTyp, qual))

			if sig.Results().Len() != 1 {
				return fmt.Errorf("function '%s' must have only 1 return value", name)
			}

			f.Comment(types.ObjectString(funcTyp, qual))
			f.Func().Id(funcTyp.Name()).
				ParamsFunc(jen.AddParams(sig.Params(), sig.Variadic(), customNamedType)).
				ParamsFunc(jen.AddParams(sig.Results(), false, customNamedType)).
				// ParamsFunc(func(rgroup *jen2.Group) {
				// 	sigParam := sig.Results().At(0)
				// 	rgroup.Id(sigParam.Name()).Add(jen.GetQualCode(sigParam.Type(), customNamedType))
				// }).
				Block(
					jen2.Return(
						jen2.Qual(ispkg, funcTyp.Name()).
							Types(jen2.Qual(sdialecttagpkg, sname+"Tag")).
							CallFunc(func(pgroup *jen2.Group) {
								for k := 0; k < sig.Params().Len(); k++ {
									sigParam := sig.Params().At(k)
									c := jen2.Id(jen.ParamName(k, sigParam))
									if sig.Variadic() && k == sig.Params().Len()-1 {
										c.Op("...")
									}
									pgroup.Add(c)
								}
							}),
					),
				)

			f.Line()
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
