package main

import "encoding/json"

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

// SliceToString Convert slice to string.
func SliceToString(input []interface{}) []byte {
	if b, err := json.Marshal(input); err == nil {
		return b
	}
	return []byte{}
}
