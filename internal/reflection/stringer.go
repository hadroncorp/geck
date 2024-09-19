package reflection

import "reflect"

func NewTypeName[T any]() string {
	var zeroVal T
	typeOf := reflect.TypeOf(zeroVal)
	return typeOf.Name()
}

func NewTypeFullName[T any]() string {
	var zeroVal T
	typeOf := reflect.TypeOf(zeroVal)
	return typeOf.String()
}

func NewTypeNameAny(v any) string {
	typeOf := reflect.TypeOf(v)
	return typeOf.Name()
}

func NewTypeFullNameAny(v any) string {
	typeOf := reflect.TypeOf(v)
	return typeOf.String()
}
