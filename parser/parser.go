package parser

import (
	"C"
	"fmt"
	"reflect"
	"regexp"
	"unsafe"
)
import "errors"

var (
	cTypeInt    = regexp.MustCompile("_Ctype_int")
	cTypeFloat  = regexp.MustCompile("_Ctype_float")
	cTypeDouble = regexp.MustCompile("_Ctype_double")
	cTypeChar   = regexp.MustCompile("_Ctype_char")
)

// CStructToGoStruct send C struct fields to Go struct fields
func CStructToGoStruct(from interface{}, to interface{}) error {
	fromValueElem := reflect.ValueOf(from)
	toValueElem := reflect.ValueOf(to)
	// toValue := reflect.New(reflect.TypeOf(to)).Elem() //Without ptr address

	if fromValueElem.Kind() != reflect.Ptr || toValueElem.Kind() != reflect.Ptr {
		return errors.New("Please pass the parameter pointers 'from' and 'to'")
	}

	fromValueElem = fromValueElem.Elem()
	toValueElem = toValueElem.Elem()

	fromValueElemType := fromValueElem.Type()
	toValueElemType := toValueElem.Type()

	var (
		fromFieldValue  reflect.Value
		fromFieldStruct reflect.StructField
		toFieldValue    reflect.Value
		toFieldStruct   reflect.StructField
		fieldExists     bool
	)

	for i := 0; i < toValueElem.NumField(); i++ {
		toFieldValue = toValueElem.Field(i)
		toFieldStruct = toValueElemType.Field(i)

		fromFieldValue = fromValueElem.FieldByName(toFieldStruct.Name)
		fromFieldStruct, fieldExists = fromValueElemType.FieldByName(toFieldStruct.Name)

		if !fieldExists {
			return fmt.Errorf("Field %s does not exist in structure C", toFieldStruct.Name)
		}

		if fromFieldStruct.Type.Kind() == reflect.Ptr || fromFieldStruct.Type.Kind() == reflect.Interface {
			fromFieldValue = fromFieldValue.Elem()
		}

		// switch fromFieldValue.Interface().(type) {
		// case C.int:
		// 	toFieldValue.SetInt(int64(fromFieldValue.Interface().(C.int)))
		// case C.float:
		// 	toFieldValue.SetFloat(float64(fromFieldValue.Interface().(C.float)))
		// case C.double:
		// 	toFieldValue.SetFloat(float64(fromFieldValue.Interface().(C.double)))
		// case C.char:
		// 	toFieldValue.SetString(C.GoString(fromFieldValue.Addr().Interface().(*C.char)))
		// default:
		// 	return fmt.Errorf("Unrecognized value type: %s", fromFieldValue.Type().String())
		// }

		Ctype := fromFieldValue.Type().String()
		// fmt.Println("Field name:", fromValueElemType.Field(i).Name, "Value:", fromFieldValue.Interface(), "Type", tStr, "Kind", fromFieldValue.Kind(), "Index:", i)

		if cTypeInt.MatchString(Ctype) {
			toFieldValue.SetInt(int64(*(*C.int)(unsafe.Pointer(fromFieldValue.Addr().Pointer()))))
		} else if cTypeFloat.MatchString(Ctype) {
			toFieldValue.SetFloat(float64(*(*C.float)(unsafe.Pointer(fromFieldValue.Addr().Pointer()))))
		} else if cTypeDouble.MatchString(Ctype) {
			toFieldValue.SetFloat(float64(*(*C.double)(unsafe.Pointer(fromFieldValue.Addr().Pointer()))))
		} else if cTypeChar.MatchString(Ctype) {
			toFieldValue.SetString(C.GoString((*C.char)(unsafe.Pointer(fromFieldValue.Addr().Pointer()))))
		} else {
			return fmt.Errorf("Unrecognized value type: %s", Ctype)
		}
	}

	return nil
}
