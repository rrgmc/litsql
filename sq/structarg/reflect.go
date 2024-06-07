package structarg

import (
	"reflect"
	"strings"
)

func (s *argValues) getStructFieldByName(value reflect.Value, fieldName string) (any, bool) {
	typ := value.Type()
	for i := 0; i < typ.NumField(); i++ {
		// Get the StructField first since this is a cheap operation. If the
		// field is unexported, then ignore it.
		f := typ.Field(i)
		if f.PkgPath != "" {
			continue
		}

		// Next get the actual value of this field and verify it is assignable
		// to the map value.
		v := value.Field(i)

		tagValue := f.Tag.Get(s.tagName)
		keyName := f.Name

		if f.Anonymous && reflect.Indirect(v).Kind() == reflect.Struct {
			// embedded struct
			eval, ok := s.getStructFieldByName(reflect.Indirect(v), fieldName)
			if ok {
				return eval, true
			}
		} else {
			// Determine the name of the key in the map
			if index := strings.Index(tagValue, ","); index != -1 {
				if tagValue[:index] == "-" {
					continue
				}

				if keyNameTagValue := tagValue[:index]; keyNameTagValue != "" {
					keyName = keyNameTagValue
				}
			} else if len(tagValue) > 0 {
				if tagValue == "-" {
					continue
				}
				keyName = tagValue
			}

			if s.mapperFunc != nil {
				keyName = s.mapperFunc(keyName)
			}

			if keyName != fieldName {
				continue
			}

			if v.Kind() == reflect.Ptr && v.IsNil() {
				// avoid sending a pointer to a nil
				return nil, true
			}
			if s.derefPointer {
				return deref(v), true
			}
			return v.Interface(), true
		}
	}

	return nil, false
}

func getReflectValue(value any) reflect.Value {
	v := reflect.ValueOf(value)
	v = reflect.Indirect(v)
	if k := v.Kind(); k != reflect.Struct {
		return reflect.Value{}
	}
	return v
}

func deref(v reflect.Value) any {
	for {
		if v.Kind() == reflect.Ptr {
			if v.IsNil() {
				return nil
			}
			v = v.Elem()
		} else {
			return v.Interface()
		}
	}
}
