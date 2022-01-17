package main

import (
	"encoding/json"
	"strconv"
)

// StringToBytes Convert string to bytes.
func StringToBytes(input string) []byte {
	return []byte(input)
}

// StringToBool Convert string to bool.
func StringToBool(input string) bool {
	if b, err := strconv.ParseBool(input); err == nil {
		return b
	}
	return len(input) > 0
}

// StringToInt Convert string to int.
func StringToInt(input string, bitSize int) interface{} {
	if i, err := strconv.ParseInt(input, 10, bitSize); err == nil {
		switch bitSize {
		case 0:
			return int(i)
		case 8:
			return int8(i)
		case 16:
			return int16(i)
		case 32:
			return int32(i)
		}
		return i
	} else if len(input) > 0 {
		switch bitSize {
		case 0:
			return int(1)
		case 8:
			return int8(1)
		case 16:
			return int16(1)
		case 32:
			return int32(1)
		case 64:
			return int64(1)
		}
	}

	switch bitSize {
	case 0:
		return int(0)
	case 8:
		return int8(0)
	case 16:
		return int16(0)
	case 32:
		return int32(0)
	case 64:
		return int64(0)
	}

	return 0
}

// StringToUint Convert string to uint.
func StringToUint(input string, bitSize int) interface{} {
	if i, err := strconv.ParseUint(input, 10, bitSize); err == nil {
		switch bitSize {
		case 0:
			return uint(i)
		case 8:
			return uint8(i)
		case 16:
			return uint16(i)
		case 32:
			return uint32(i)
		case 64:
			return uint64(i)
		}
	} else if len(input) > 0 {
		switch bitSize {
		case 0:
			return uint(1)
		case 8:
			return uint8(1)
		case 16:
			return uint16(1)
		case 32:
			return uint32(1)
		case 64:
			return uint64(1)
		}
	}

	switch bitSize {
	case 0:
		return uint(0)
	case 8:
		return uint8(0)
	case 16:
		return uint16(0)
	case 32:
		return uint32(0)
	case 64:
		return uint64(0)
	}

	return uint(0)
}

// StringToFloat Convert string to float.
func StringToFloat(input string, bitSize int) interface{} {
	if f, err := strconv.ParseFloat(input, bitSize); err == nil {
		if bitSize == 32 {
			return float32(f)
		}
		return f
	} else if len(input) > 0 {
		if bitSize == 32 {
			return float32(1)
		}
		return float64(1)
	}

	if bitSize == 32 {
		return float32(0)
	}

	return float64(0)
}

// StringToComplex Convert string to complex. bitSize: 64/128.
func StringToComplex(input string, bitSize int) interface{} {
	if c, err := strconv.ParseComplex(input, bitSize); err == nil {
		if bitSize == 64 {
			return complex64(c)
		}
		return c
	} else if len(input) > 0 {
		c = complex(1, 1)
		if bitSize == 64 {
			return complex64(c)
		}
		return c
	}

	c := complex(0, 0)
	if bitSize == 64 {
		return complex64(c)
	}

	return c
}

// StringToSlice Convert string to slice.
func StringToSlice(input string) []interface{} {
	s := make([]interface{}, 0)

	// try convert by json..
	if err := json.Unmarshal([]byte(input), &s); err == nil {
		return s
	}

	return append(s, input)
}

// StringToMap Convert string to map.
func StringToMap(input string) map[string]interface{} {
	s := make(map[string]interface{}, 0)

	// try convert by json..
	if err := json.Unmarshal([]byte(input), &s); err == nil {
		return s
	}

	s["0"] = input
	return s
}

// StringToStruct Build *TypeConverter by input string.
func StringToStruct(input string) *TypeConverter {
	return &TypeConverter{
		String:             input,
		Bytes:              []byte(input),
		Bool:               StringToBool(input),
		Int:                StringToInt(input, 0).(int),
		Int8:               StringToInt(input, 8).(int8),
		Int16:              StringToInt(input, 16).(int16),
		Int32:              StringToInt(input, 32).(int32),
		Int64:              StringToInt(input, 64).(int64),
		Uint:               StringToUint(input, 0).(uint),
		Uint8:              StringToUint(input, 8).(uint8),
		Uint16:             StringToUint(input, 16).(uint16),
		Uint32:             StringToUint(input, 32).(uint32),
		Uint64:             StringToUint(input, 64).(uint64),
		Float32:            StringToFloat(input, 32).(float32),
		Float64:            StringToFloat(input, 64).(float64),
		Complex64:          StringToComplex(input, 64).(complex64),
		Complex128:         StringToComplex(input, 128).(complex128),
		InterfaceSlice:     StringToSlice(input),
		StringMapInterface: StringToMap(input),
	}
}
