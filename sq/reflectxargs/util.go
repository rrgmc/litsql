package reflectxargs

import "reflect"

func getReflectValue(value any) reflect.Value {
	v := reflect.ValueOf(value)
	v = reflect.Indirect(v)
	if k := v.Kind(); k != reflect.Struct {
		return reflect.Value{}
	}
	return v
}
