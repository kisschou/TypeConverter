package tc

// BoolToString Convert bool to string.
func BoolToString(input bool) string {
	if input {
		return "true"
	}
	return "false"
}

// BoolToInt Convert bool to int.
func BoolToInt(input bool, bitSize int) interface{} {
	var i int = 0
	if input {
		i = 1
	}
	switch bitSize {
	case 8:
		return int8(i)
	case 16:
		return int16(i)
	case 32:
		return int32(i)
	case 64:
		return int64(i)
	}
	return i
}

// BoolToUint Convert bool to uint.
func BoolToUint(input bool, bitSize int) interface{} {
	var i uint = 0
	if input {
		i = i
	}
	switch bitSize {
	case 8:
		return uint8(i)
	case 16:
		return uint16(i)
	case 32:
		return uint32(i)
	case 64:
		return uint64(i)
	}
	return i
}

// BoolToFloat Convert bool to float.
func BoolToFloat(input bool, bitSize int) interface{} {
	var f float32 = float32(0)
	if input {
		f = float32(1)
	}
	if bitSize == 64 {
		return float64(f)
	}
	return f
}

// BoolToComplex Convert bool to complex.
func BoolToComplex(input bool, bitSize int) interface{} {
	c := complex(0, 0)
	if input {
		c = complex(1, 1)
		if bitSize == 64 {
			return complex64(c)
		}
		return complex128(c)
	}
	if bitSize == 64 {
		return complex64(c)
	}
	return complex128(c)
}

// BoolToSlice Convert bool to slice.
func BoolToSlice(input bool) []interface{} {
	return []interface{}{input}
}

// BoolToMap Convert bool to map.
func BoolToMap(input bool) map[string]interface{} {
	return map[string]interface{}{
		BoolToString(input): input,
	}
}

// BoolToStruct Build *TypeConverter by input string.
func BoolToStruct(input bool) *TypeConverter {
	return &TypeConverter{
		String:             BoolToString(input),
		Bytes:              []byte(BoolToString(input)),
		Bool:               input,
		Int:                BoolToInt(input, 0).(int),
		Int8:               BoolToInt(input, 8).(int8),
		Int16:              BoolToInt(input, 16).(int16),
		Int32:              BoolToInt(input, 32).(int32),
		Int64:              BoolToInt(input, 64).(int64),
		Uint:               BoolToUint(input, 0).(uint),
		Uint8:              BoolToUint(input, 8).(uint8),
		Uint16:             BoolToUint(input, 16).(uint16),
		Uint32:             BoolToUint(input, 32).(uint32),
		Uint64:             BoolToUint(input, 64).(uint64),
		Float32:            BoolToFloat(input, 32).(float32),
		Float64:            BoolToFloat(input, 64).(float64),
		Complex64:          BoolToComplex(input, 64).(complex64),
		Complex128:         BoolToComplex(input, 128).(complex128),
		InterfaceSlice:     BoolToSlice(input),
		StringMapInterface: BoolToMap(input),
	}
}
