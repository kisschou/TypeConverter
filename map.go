package main

import "encoding/json"

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

// MapToString Convert map to string.
func MapToString(input map[string]interface{}) []byte {
	if b, err := json.Marshal(input); err == nil {
		return b
	}
	return []byte{}
}
