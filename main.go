package main

import (
	"encoding/json"
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

// ====> String To

// StringToBytes Convert string to bytes.
func StringToBytes(input string) []byte {
	return []byte(input)
}

// StringToBool Convert string to bool.
func StringToBool(input string) bool {
	return len(input) > 0
}

// StringToInt Convert string to int.
func StringToInt(input string, bitSize int) interface{} {
	// to int.
	if bitSize == 0 {
		if i, err := strconv.Atoi(input); err == nil {
			return i
		}
		return 0
	}

	// to other..
	if i, err := strconv.ParseInt(input, 10, bitSize); err == nil {
		return i
	}

	return 0
}

// StringToFloat Convert string to float.
func StringToFloat(input string, bitSize int) interface{} {
	if f, err := strconv.ParseFloat(input, bitSize); err == nil {
		return f
	}
	return 0
}

// StringToSlice Convert string to slice.
func StringToSlice(input string) []interface{} {
	return []interface{}{input}
}

// StringToMap Convert string to map.
func StringToMap(input string) map[string]interface{} {
	return map[string]interface{}{"0": input}
}

// <==== End

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

// SliceToString Convert slice to string.
func SliceToString(input []interface{}) []byte {
	if b, err := json.Marshal(input); err == nil {
		return b
	}
	return []byte{}
}

// MapToString Convert map to string.
func MapToString(input map[string]interface{}) []byte {
	if b, err := json.Marshal(input); err == nil {
		return b
	}
	return []byte{}
}

// SliceToInterfaceSlice Convert slice to []interface{}.
func SliceToInterfaceSlice(input interface{}) []interface{} {
	r := make([]interface{}, 0)
	/*
		for k, v := range reflect.ValueOf(input) {
			r[k] = v
		}
	*/
	return r
}

// MapToStringMapInterface Convert map to map[string]interface{}.
func MapToStringMapInterface(input interface{}) map[string]interface{} {
	r := make(map[string]interface{}, 0)
	/*
		for k, v := range input {
			r[k] = v
		}
	*/
	return r
}

// NewConverter Init *TypeConverter.
func NewConverter(input interface{}) *TypeConverter {
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
	var (
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
	)

	switch tName {
	case "string":
		String = input.(string)
		Bytes = []byte(String)
		Bool = StringToBool(String)
		Int = StringToInt(String, 0)
		Int8 = StringToInt(String, 8)
		Int16 = StringToInt(String, 16)
		Int32 = StringToInt(String, 32)
		Int64 = StringToInt(String, 64)
		Uint = uint(Int)
		Uint8 = uint8(Int8)
		Uint16 = uint16(Int16)
		Uint32 = uint32(Int32)
		Uint64 = uint64(Int64)
		InterfaceSlice = StringToSlice(String)
		StringMapInterface = StringToMap(String)
		break

	case "bool":
		Bool = input.(bool)
		String = BoolToString(Bool)
		Bytes = []byte(String)
		Int = BoolToInt(Bool)
		InterfaceSlice = StringToSlice(String)
		StringMapInterface = StringToMap(String)
		break

	case "byte":
		Byte = input.(byte)
		Bytes = []byte{Byte}
		String = string(Bytes)
		Bool = StringToBool(String)
		Int = StringToInt(String)
		InterfaceSlice = SliceToInterfaceSlice(Bytes)
		StringMapInterface = StringToMap(String)
		break

	case "int":
		Int = input.(int)
		String = strconv.Itoa(Int)
		Bytes = []byte(String)
		Bool = StringToBool(String)
		InterfaceSlice = StringToSlice(String)
		StringMapInterface = StringToMap(String)
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

	Float32 = float32(Int)
	Float64 = float64(Int)

	return &TypeConverter{
		String:             String,
		Byte:               Byte,
		Bytes:              Bytes,
		Bool:               Bool,
		Int:                Int,
		Int8:               Int8,
		Int16:              Int16,
		Int32:              Int32,
		Int64:              Int64,
		Uint:               Uint,
		Uint8:              Uint8,
		Uint16:             Uint16,
		Uint32:             Uint32,
		Uint64:             Uint64,
		Float32:            Float32,
		Float64:            Float64,
		Complex64:          Complex64,
		Complex128:         Complex128,
		InterfaceSlice:     InterfaceSlice,
		StringMapInterface: StringMapInterface,
	}
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
