package genutil

import (
	"fmt"
	"go/types"
)

func ParamName(idx int, param *types.Var) string {
	if param.Name() != "" {
		return param.Name()
	}
	return fmt.Sprintf("p%d", idx)
}

func FormatObjectName(obj types.Object) string {
	pkg := ""
	if obj.Pkg().Name() != "" {
		pkg += obj.Pkg().Name()
	}
	return pkg + "." + obj.Name()
}
