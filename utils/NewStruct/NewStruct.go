package NewStruct

import "reflect"

func RecursiveNewStruct2[out any](in *out) *out {
	val := reflect.ValueOf(in)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		if fieldType.Type.Kind() == reflect.Struct {
			// If the field is a struct, recursively initialize it
			field.Set(
				reflect.ValueOf(RecursiveNewStruct(reflect.New(field.Type()).Interface())).
					Elem(),
			)
		} else {
			// Otherwise, set a default value based on the field type
			zero := reflect.Zero(fieldType.Type)
			field.Set(zero)
		}
	}

	return in
}

func RecursiveNewStruct(obj interface{}) interface{} {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		if fieldType.Type.Kind() == reflect.Struct {
			// If the field is a struct, recursively initialize it
			field.Set(
				reflect.ValueOf(RecursiveNewStruct(reflect.New(field.Type()).Interface())).
					Elem(),
			)
		} else {
			// Otherwise, set a default value based on the field type
			zero := reflect.Zero(fieldType.Type)
			field.Set(zero)
		}
	}

	return val.Addr().Interface()
}
