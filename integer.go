package tc

import (
	"strconv"
)

// IntToString Convert int to string.
func IntToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// UintToString Convert uint to string.
func UintToString(i uint64) string {
	return strconv.FormatUint(i, 10)
}

// FloatToString Convert float64 to string.
func FloatToString(i float64) string {
	return strconv.FormatFloat(i, 'E', -1, 64)
}

// ComplexToString Convert complex128 to string.
func ComplexToString(c complex128) string {
	return strconv.FormatComplex(c, 'E', -1, 128)
}
