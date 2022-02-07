package tc

import (
	"encoding/json"
	"reflect"
)

// MapToStringMapInterface Convert map[string]interface{} to map[string]interface{}.
func MapToStringMapInterface(input interface{}) map[string]interface{} {
	r := make(map[string]interface{}, 0)
	m := reflect.ValueOf(input)
	keys := m.MapKeys()
	for _, key := range keys {
		r[key.Interface().(string)] = m.MapIndex(key).Interface()
	}
	return r
}

// MapToString Convert map[string]interface{} to []byte.
func MapToBytes(input map[string]interface{}) []byte {
	if b, err := json.Marshal(input); err == nil {
		return b
	}
	return []byte{}
}

// MapToSlice Convert map[string]interface{} to []interface{}.
func MapToSlice(input map[string]interface{}) []interface{} {
	s := make([]interface{}, 0)
	for k, v := range input {
		s = append(s, k)
		s = append(s, v)
	}
	return s
}

// MapToStruct Convert map[string]interface{} to struct.
func MapToStruct(input map[string]interface{}) *TypeConverter {
	boolVal := len(input) > 0
	return &TypeConverter{
		String:     string(MapToBytes(input)),
		Bytes:      MapToBytes(input),
		Bool:       boolVal,
		Int:        BoolToInt(boolVal, 0).(int),
		Int8:       BoolToInt(boolVal, 8).(int8),
		Int16:      BoolToInt(boolVal, 16).(int16),
		Int32:      BoolToInt(boolVal, 32).(int32),
		Int64:      BoolToInt(boolVal, 64).(int64),
		Uint:       BoolToUint(boolVal, 0).(uint),
		Uint8:      BoolToUint(boolVal, 8).(uint8),
		Uint16:     BoolToUint(boolVal, 16).(uint16),
		Uint32:     BoolToUint(boolVal, 32).(uint32),
		Uint64:     BoolToUint(boolVal, 64).(uint64),
		Float32:    BoolToFloat(boolVal, 32).(float32),
		Float64:    BoolToFloat(boolVal, 64).(float64),
		Complex64:  BoolToComplex(boolVal, 64).(complex64),
		Complex128: BoolToComplex(boolVal, 128).(complex128),
		Slice:      MapToSlice(input),
		Map:        input,
	}
}
