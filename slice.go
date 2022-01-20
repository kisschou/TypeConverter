package tc

import (
	"encoding/json"
	"reflect"
	"strconv"
)

// SliceToInterfaceSlice Convert slice to []interface{}.
func SliceToInterfaceSlice(input interface{}) []interface{} {
	l := make([]interface{}, 0)
	s := reflect.ValueOf(input)
	for i := 0; i < s.Len(); i++ {
		l = append(l, s.Index(i).Interface())
	}
	return l
}

// SliceIsBytes check is []uint8 or []byte.
func SliceIsBytes(input []interface{}) bool {
	for _, v := range input {
		tName := reflect.TypeOf(v).String()
		if tName != "uint8" && tName != "byte" {
			return false
		}
	}
	return true
}

// SliceToString Convert slice to string.
func SliceToString(input []interface{}) string {
	b := SliceToBytes(input)
	return string(b)
}

// SliceToString Convert slice to string.
func SliceToBytes(input []interface{}) []byte {
	b := make([]byte, 0)
	if SliceIsBytes(input) {
		for _, v := range input {
			if reflect.TypeOf(v).String() == "uint8" {
				b = append(b, byte(v.(uint8)))
			} else {
				b = append(b, v.(byte))
			}
		}
	} else {
		if b, err := json.Marshal(input); err == nil {
			return b
		}
		b = make([]byte, 0)
	}
	return b
}

// SliceToMap Convert slice to map[string]interface{}
func SliceToMap(input []interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	for k, v := range input {
		m[strconv.Itoa(k)] = v
	}
	return m
}

// SliceToStruct Convert slice to struct.
func SliceToStruct(input []interface{}) *TypeConverter {
	boolVal := len(input) > 0
	return &TypeConverter{
		String:             SliceToString(input),
		Bytes:              SliceToBytes(input),
		Bool:               boolVal,
		Int:                BoolToInt(boolVal, 0).(int),
		Int8:               BoolToInt(boolVal, 8).(int8),
		Int16:              BoolToInt(boolVal, 16).(int16),
		Int32:              BoolToInt(boolVal, 32).(int32),
		Int64:              BoolToInt(boolVal, 64).(int64),
		Uint:               BoolToUint(boolVal, 0).(uint),
		Uint8:              BoolToUint(boolVal, 8).(uint8),
		Uint16:             BoolToUint(boolVal, 16).(uint16),
		Uint32:             BoolToUint(boolVal, 32).(uint32),
		Uint64:             BoolToUint(boolVal, 64).(uint64),
		Float32:            BoolToFloat(boolVal, 32).(float32),
		Float64:            BoolToFloat(boolVal, 64).(float64),
		Complex64:          BoolToComplex(boolVal, 64).(complex64),
		Complex128:         BoolToComplex(boolVal, 128).(complex128),
		InterfaceSlice:     input,
		StringMapInterface: SliceToMap(input),
	}
}
