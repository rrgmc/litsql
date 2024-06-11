package genutil

import (
	"fmt"
	"go/types"
	"strings"

	"github.com/dave/jennifer/jen"
)

type CustomType interface {
	NamedType(st jen.Statement, tt *types.Named) *jen.Statement
	IsPointer(tt *types.Pointer) bool
}

// type CustomType func(st jen.Statement, tt *types.Named) *jen.Statement

func GetQualCode(typ types.Type, customType CustomType) *jen.Statement {
	var st jen.Statement
	for {
		switch tt := typ.(type) {
		case *types.Basic:
			return st.Add(jen.Id(tt.Name()))
		case *types.Array:
			return st.Add(jen.Index(jen.Lit(tt.Len())).Add(GetQualCode(tt.Elem(), customType)))
		case *types.Slice:
			return st.Add(jen.Index().Add(GetQualCode(tt.Elem(), customType)))
		case *types.Pointer:
			if customType == nil || customType.IsPointer(tt) {
				st.Add(jen.Op("*"))
			}
			typ = tt.Elem()
		case *types.Tuple:
			var items jen.Statement
			for i := 0; i < tt.Len(); i++ {
				items.Add(jen.Id(tt.At(i).Name()).Add(GetQualCode(tt.At(i).Type(), customType)))
			}
			return st.Add(jen.Params(items...))
		case *types.Interface:
			return st.Add(jen.Id(tt.String()))
		case *types.Map:
			return st.Add(jen.Map(GetQualCode(tt.Key(), customType)).Add(GetQualCode(tt.Elem(), customType)))
		case *types.Chan:
			var chanDesc *jen.Statement
			switch tt.Dir() {
			case types.SendRecv:
				chanDesc = jen.Chan()
			case types.SendOnly:
				chanDesc = jen.Chan().Op("<-")
			case types.RecvOnly:
				chanDesc = jen.Op("<-").Chan()
			default:
				panic("unknown channel direction")
			}
			return st.Add(chanDesc.Add(GetQualCode(tt.Elem(), customType)))
		case *types.Named:
			if tt.Obj().Pkg() != nil {
				if customType != nil {
					customRet := customType.NamedType(st, tt)
					if customRet != nil {
						return customRet
					}
				}
				return st.Add(jen.Qual(tt.Obj().Pkg().Path(), tt.Obj().Name()).TypesFunc(AddTypeList(tt.TypeArgs(), customType)))
			}
			return st.Add(jen.Id(tt.Obj().Name()).TypesFunc(AddTypeList(tt.TypeArgs(), customType)))
		case *types.TypeParam:
			return st.Add(jen.Id(tt.Obj().Name()))
		case *types.Signature:
			return st.Add(jen.Func().
				ParamsFunc(AddParams(tt.Params(), tt.Variadic(), customType)).
				ParamsFunc(AddParams(tt.Results(), false, customType)))
		default:
			panic(fmt.Errorf("unknown type %T", typ))
		}
	}
}

func TypeNameCode(typeName string) (*jen.Statement, error) {
	typeName, isPtr := strings.CutPrefix(typeName, "*")

	var st jen.Statement
	if isPtr {
		st.Add(jen.Op("*"))
	}

	lastIndex := strings.LastIndexAny(typeName, "/.")
	if lastIndex == -1 {
		return st.Add(jen.Id(typeName)), nil
	}
	if typeName[lastIndex:lastIndex+1] == "." {
		return st.Add(jen.Qual(typeName[:lastIndex], typeName[lastIndex+1:])), nil
	}
	return nil, fmt.Errorf("invalid type name format (must have a dot to determine the type): %s", typeName)
}

func AddTypeParamsList(typeList *types.TypeParamList, withType bool, customNamedType CustomType) func(*jen.Group) {
	return func(tgroup *jen.Group) {
		for t := 0; t < typeList.Len(); t++ {
			tparam := typeList.At(t)
			if withType {
				tgroup.Id(tparam.Obj().Name()).Add(GetQualCode(tparam.Constraint(), customNamedType))
			} else {
				tgroup.Id(tparam.Obj().Name())
			}
		}
	}
}

func AddTypeList(typeList *types.TypeList, customNamedType CustomType) func(*jen.Group) {
	return func(tgroup *jen.Group) {
		for t := 0; t < typeList.Len(); t++ {
			tparam := typeList.At(t)
			tgroup.Add(GetQualCode(tparam, customNamedType))
		}
	}
}

func AddParams(params *types.Tuple, isVariadic bool, customNamedType CustomType) func(*jen.Group) {
	return func(group *jen.Group) {
		for k := 0; k < params.Len(); k++ {
			sigParam := params.At(k)
			c := &jen.Statement{}
			if sigParam.Name() != "" {
				c.Add(jen.Id(ParamName(k, sigParam)))
			}
			if isVariadic && k == params.Len()-1 {
				c.Add(jen.Op("...").Add(GetQualCode(sigParam.Type().(*types.Slice).Elem(), customNamedType)))
			} else {
				c.Add(GetQualCode(sigParam.Type(), customNamedType))
			}
			group.Add(c)
		}
	}
}

type CTFunc struct {
	FuncNamedType func(st jen.Statement, tt *types.Named) *jen.Statement
	FuncIsPointer func(tt *types.Pointer) bool
}

func (f *CTFunc) NamedType(st jen.Statement, tt *types.Named) *jen.Statement {
	return f.FuncNamedType(st, tt)
}

func (f *CTFunc) IsPointer(tt *types.Pointer) bool {
	return f.FuncIsPointer(tt)
}
