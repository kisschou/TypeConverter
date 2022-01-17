package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type (
	TypeConverter struct {
		// String string
		String string

		// Byte byte
		Byte byte

		// Bytes []byte
		Bytes []byte

		// Bool bool
		Bool bool

		// Int int
		Int int

		// Int8 int8
		Int8 int8

		// Int16 int16
		Int16 int16

		// Int32 int32
		Int32 int32

		// Int64 int64
		Int64 int64

		// Uint unsigened int
		Uint uint

		// Uint8 unsigened int8
		Uint8 uint8

		// Uint16 unsigened int16
		Uint16 uint16

		// Uint32 unsigened int32
		Uint32 uint32

		// Uint64 unsigened int64
		Uint64 uint64

		// Float32 float32
		Float32 float32

		// Float64 float64
		Float64 float64

		// Complex64 complex64
		Complex64 complex64

		// Complex128 complex128
		Complex128 complex128

		// InterfaceSlice []interface{}
		InterfaceSlice []interface{}

		// StringMapInterface map[string]interface{}
		StringMapInterface map[string]interface{}
	}
)

// ====> Bool To

// BoolToString Convert bool to string.
func BoolToString(input bool) string {
	if input {
		return "true"
	}
	return "false"
}

// BoolToInt Convert bool to int.
func BoolToInt(input bool) int {
	if input {
		return 1
	}
	return 0
}

// <==== End

// ====> Int to

// IntToString Convert int to string.
func IntToString(i int) string {
	return strconv.Itoa(i)
}

// <====

// NewConverter Init *TypeConverter.
func NewConverter(input interface{}) *TypeConverter {
	tStruct := new(TypeConverter)
	tName := reflect.TypeOf(input).String()
	/*
		fmt.Println("=====================>")
		fmt.Println("TypeOf: ", reflect.TypeOf(input))
		fmt.Println("TypeOf Name(): ", reflect.TypeOf(input).Name())
		fmt.Println("TypeOf String(): ", reflect.TypeOf(input).String())
		fmt.Println("TypeOf Kind().String(): ", reflect.TypeOf(input).Kind().String())
		fmt.Println("TypeOf Bits(): ", reflect.TypeOf(input).Kind().String())
		fmt.Println("TypeOf Comparable(): ", reflect.TypeOf(input).Comparable())
		fmt.Println("<=====================")
	*/

	switch tName {
	case "string":
		tStruct = StringToStruct(input.(string))
		break

	case "bool":
		Bool = input.(bool)
		break

	case "byte":
		Byte = input.(byte)
		break

	case "int":
		Int = input.(int)
		break
	case "int8":
		Int8 = input.(int8)
		break
	case "int16":
		Int16 = input.(int16)
		break
	case "int32":
		Int32 = input.(int32)
		break
	case "int64":
		Int64 = input.(int64)
		break
	case "uint":
		Uint = input.(uint)
		break
	case "uint8":
		Uint8 = input.(uint8)
		break
	case "uint16":
		Uint16 = input.(uint16)
		break
	case "uint32":
		Uint32 = input.(uint32)
		break
	case "uint64":
		Uint64 = input.(uint64)
		break
	case "float32":
		Float32 = input.(float32)
		break
	case "float64":
		Float64 = input.(float64)
		break

	case "complex64":
		Complex64 = input.(complex64)
		break
	case "complex128":
		Complex128 = input.(complex128)
		break

	case "[]interface{}":
		break

	case "map[string]interface{}":
		break
	}

	return tStruct
}

func main() {
	/*
		NewConverter([]map[string]interface{}{
			map[string]interface{}{"a": 1, "b": "b"},
		})
		return
	*/
	rawData := "3306"
	nc := NewConverter(rawData)
	fmt.Println(nc)
	return

	fmt.Println(fmt.Sprintf("pool_size => %v", rawData), 14)
	fmt.Println(fmt.Sprintf("pool_size => %v", NewConverter(rawData)), 14)
	fmt.Println("ValueOf: ", reflect.ValueOf(rawData))
	fmt.Println("TypeOf: ", reflect.TypeOf(rawData))
	fmt.Println("TypeOf Name(): ", reflect.TypeOf(rawData).Name())
	fmt.Println("TypeOf String(): ", reflect.TypeOf(rawData).String())
	fmt.Println("TypeOf Kind().String(): ", reflect.TypeOf(rawData).Kind().String())
	fmt.Println("TypeOf Bits(): ", reflect.TypeOf(rawData).Kind().String())
	fmt.Println("TypeOf Comparable(): ", reflect.TypeOf(rawData).Comparable())
}
