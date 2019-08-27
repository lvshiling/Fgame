package reflectutils

import (
	"fmt"
	"reflect"
	"strconv"
)

type InvalidPrimitiveError struct {
	typ reflect.Type
}

func (ipe *InvalidPrimitiveError) Error() string {
	return fmt.Sprintf("invalid primitive type (%s)", ipe.typ.String())
}

// func IsPrimitive(typ reflect.Type) bool {
// 	switch typ.Kind() {
// 	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
// 	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
// 	case reflect.Bool:
// 	case reflect.Float32:
// 	case reflect.Float64:
// 	case reflect.String:
// 		return true
// 	}
// 	return false
// }

//todo:improve complex
//except complex64 and complex 128
func ParsePrimitive(typ reflect.Type, val string) (interface{}, error) {

	tvalue := reflect.New(typ)
	value := tvalue.Elem()
	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

		intVal, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return nil, err
		}
		value.SetInt(intVal)
		return value.Interface(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:

		uintVal, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			return nil, err
		}
		value.SetUint(uintVal)
		return value.Interface(), nil
	case reflect.Bool:

		boolVal, err := strconv.ParseBool(val)
		if err != nil {
			return nil, err
		}

		value.SetBool(boolVal)
		return value.Interface(), nil
	case reflect.Float32:

		floatVal, err := strconv.ParseFloat(val, 32)
		if err != nil {
			return nil, err
		}
		value.SetFloat(floatVal)
		return value.Interface(), nil
	case reflect.Float64:

		floatVal, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return nil, err
		}
		value.SetFloat(floatVal)
		return value.Interface(), nil
	case reflect.String:

		value.SetString(val)
		return value.Interface(), nil
	default:
		return nil, &InvalidPrimitiveError{typ}
	}
}
