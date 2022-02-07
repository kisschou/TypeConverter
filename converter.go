package tc

import (
	"reflect"
	"strings"
)

type (
	// TypeConverter .
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

		// Slice []interface{}
		Slice []interface{}

		// Map map[string]interface{}
		Map map[string]interface{}
	}
)

// New Init *TypeConverter.
func New(input interface{}) *TypeConverter {
	tStruct := new(TypeConverter)
	tKind := reflect.TypeOf(input).Kind().String()
	tValue := reflect.ValueOf(input)

	switch tKind {
	case "string":
		tStruct = StringToStruct(tValue.String())
		break
	case "bool":
		tStruct = BoolToStruct(tValue.Bool())
		break
	case "byte":
		tStruct = StringToStruct(string(input.(byte)))
		break
	case "uint8":
		tStruct = StringToStruct(string(byte(input.(uint8))))
		break
	case "int", "int8", "int16", "int32", "int64":
		tStruct = StringToStruct(IntToString(tValue.Int()))
		break
	case "uint", "uint16", "uint32", "uint64":
		tStruct = StringToStruct(UintToString(tValue.Uint()))
		break
	case "float32", "float64":
		tStruct = StringToStruct(FloatToString(tValue.Float()))
		break
	case "complex64", "complex128":
		tStruct = StringToStruct(ComplexToString(tValue.Complex()))
		break
	case "slice":
		tStruct = SliceToStruct(SliceToInterfaceSlice(input))
		break
	case "map":
		tStruct = MapToStruct(MapToStringMapInterface(input))
		break
	}

	return tStruct
}

// Equal Pass in a new object and compare it with the current object to check if the values are equal.
func (tc *TypeConverter) Equal(compareData interface{}) bool {
	ctc := New(compareData)
	// 如果完全相等
	if reflect.DeepEqual(tc, ctc) {
		return true
	}
	// 校验字符串
	if tc.String == ctc.String {
		return true
	}
	// 校验切片
	if SliceEqual(tc.Slice, ctc.Slice) {
		return true
	}
	// 校验map
	if MapEqual(tc.Map, ctc.Map) {
		return true
	}
	return false
}

// Compare Pass in the new object and operator to compare against the current object.
// The optional operators are:
// EQ  -  =
// NEQ -  !=
// GT  -  >
// EGT -  >=
// LT  -  <
// ELT -  <=
func (tc *TypeConverter) Compare(operator string, compareData interface{}) bool {
	switch strings.ToUpper(operator) {
	case "EQ":
		return tc.Equal(compareData)
	case "NEQ":
		return !tc.Equal(compareData)
	case "GT":
		return tc.Int64 > New(compareData).Int64
	case "EGT":
		return tc.Int64 >= New(compareData).Int64
	case "LT":
		return tc.Int64 < New(compareData).Int64
	case "ELT":
		return tc.Int64 <= New(compareData).Int64
	}
	return false
}

// SliceEqual Compares two slices for equality.
func SliceEqual(d1, d2 []interface{}) bool {
	// 切片数量相同
	if len(d1) == len(d2) {
		// 数据及顺序完全一致
		if reflect.DeepEqual(d1, d2) {
			return true
		}

		// 如果只有一条数据, Game Over.
		if len(d1) == 1 {
			return false
		}

		// 不计顺序, 校验数据是否完全相等
		for _, v1 := range d1 {
			isFound := false
			for k, v2 := range d2 {
				// 如果找到相同项 // 删除当前项 // 退出子循环
				if New(v1).Equal(v2) {
					isFound = true
					d2 = append(d2[:k], d2[k+1:]...)
					break
				}
			}
			// 没有找到相同项
			if !isFound {
				return false
			}
		}
		return true
	}
	return false
}

// MapEqual Compares two maps for equality.
func MapEqual(d1, d2 map[string]interface{}) bool {
	// 数量相同
	if len(d1) == len(d2) {
		// 数据及顺序完全一致
		if reflect.DeepEqual(d1, d2) {
			return true
		}

		// 如果只有一条数据, Game Over.
		if len(d1) == 1 {
			return false
		}

		// 不计顺序, 校验数据是否完全相等
		for k, v1 := range d1 {
			if v2, ok := d2[k]; ok {
				if !New(v1).Equal(v2) {
					return false
				}
			} else {
				return false
			}
		}
		return true
	}
	return false
}
