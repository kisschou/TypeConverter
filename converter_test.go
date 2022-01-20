package tc

import (
	"fmt"
	"testing"
)

// 测试 ====>

// TestNew .
func TestNew(t *testing.T) {
	var tests = []struct {
		input interface{}
		s     string
		b     bool
		i     int64
		is    []interface{}
		smi   map[string]interface{}
	}{
		{input: "3306", s: "3306", b: true, i: int64(3306), is: []interface{}{"3306"}, smi: map[string]interface{}{"0": "3306"}},
		{input: "HelloWorld", s: "HelloWorld", b: true, i: int64(1), is: []interface{}{"HelloWorld"}, smi: map[string]interface{}{"0": "HelloWorld"}},
		{input: 3306, s: "3306", b: true, i: int64(3306), is: []interface{}{"3306"}, smi: map[string]interface{}{"0": "3306"}},
		{input: false, s: "false", b: false, i: int64(0), is: []interface{}{false}, smi: map[string]interface{}{"false": false}},
		{input: int64(3306), s: "3306", b: true, i: int64(3306), is: []interface{}{"3306"}, smi: map[string]interface{}{"0": "3306"}},
		{input: []string{"Hello", "World", "!"}, s: "[\"Hello\",\"World\",\"!\"]", b: true, i: int64(1), is: []interface{}{"Hello", "World", "!"}, smi: map[string]interface{}{"0": "Hello", "1": "World", "2": "!"}},
		{input: []int{1, 2, 3}, s: "[1,2,3]", b: true, i: int64(1), is: []interface{}{1, 2, 3}, smi: map[string]interface{}{"0": 1, "1": 2, "2": 3}},
		{input: []interface{}{"Hello", "World", 123}, s: "[\"Hello\",\"World\",123]", b: true, i: int64(1), is: []interface{}{"Hello", "World", 123}, smi: map[string]interface{}{"0": "Hello", "1": "World", "2": 123}},
		{input: map[string]string{"hi": "HelloWorld", "name": "Kisschou"}, s: "{\"hi\":\"HelloWorld\",\"name\":\"Kisschou\"}", b: true, i: int64(1), is: []interface{}{"hi", "HelloWorld", "name", "Kisschou"}, smi: map[string]interface{}{"hi": "HelloWorld", "name": "Kisschou"}},
		{input: map[string]int{"name": 1, "age": 13}, s: "{\"age\":13,\"name\":1}", b: true, i: int64(1), is: []interface{}{"name", 1, "age", 13}, smi: map[string]interface{}{"name": 1, "age": 13}},
		{input: map[string]interface{}{"name": 1, "age": 13, "hi": "HelloWorld"}, s: "{\"age\":13,\"hi\":\"HelloWorld\",\"name\":1}", b: true, i: int64(1), is: []interface{}{"name", 1, "age", 13, "hi", "HelloWorld"}, smi: map[string]interface{}{"name": 1, "age": 13, "hi": "HelloWorld"}},
		{input: map[string][]string{"s": []string{"Hi", "hi", "hI"}, "v": []string{"v1.0"}}, s: "{\"s\":[\"Hi\",\"hi\",\"hI\"],\"v\":[\"v1.0\"]}", b: true, i: int64(1), is: []interface{}{"s", []string{"Hi", "hi", "hI"}, "v", []string{"v1.0"}}, smi: map[string]interface{}{"s": []string{"Hi", "hi", "hI"}, "v": []string{"v1.0"}}},
	}

	for _, test := range tests {
		tStruct := New(test.input)
		if tStruct.String != test.s {
			t.Errorf("Fail verify string. got: %v, want: %v.", tStruct.String, test.s)
		}
		if tStruct.Bool != test.b {
			t.Errorf("Fail verify bool Fail. got: %v, want: %v.", tStruct.Bool, test.b)
		}
		if tStruct.Int64 != test.i {
			t.Errorf("Fail verify int64 Fail. got: %v, want: %v.", tStruct.Int64, test.i)
		}
		if !SliceEqual(tStruct.InterfaceSlice, test.is) {
			t.Errorf("Fail verify slice. got: %v, want: %v.", tStruct.InterfaceSlice, test.is)
		}
		if !MapEqual(tStruct.StringMapInterface, test.smi) {
			t.Errorf("Fail verify map. got: %v, want: %v.", tStruct.StringMapInterface, test.smi)
		}
	}
}

// <=========

// 基准 ====>

// BenchmarkNew .
// => go test -bench=. -benchmem
func BenchmarkNew(b *testing.B) {
	var tests = []struct {
		input interface{}
		s     string
		b     bool
		i     int64
		is    []interface{}
		smi   map[string]interface{}
	}{
		{input: "3306", s: "3306", b: true, i: int64(3306), is: []interface{}{"3306"}, smi: map[string]interface{}{"0": "3306"}},
		{input: "HelloWorld", s: "HelloWorld", b: true, i: int64(1), is: []interface{}{"HelloWorld"}, smi: map[string]interface{}{"0": "HelloWorld"}},
		{input: 3306, s: "3306", b: true, i: int64(3306), is: []interface{}{"3306"}, smi: map[string]interface{}{"0": "3306"}},
		{input: false, s: "false", b: false, i: int64(0), is: []interface{}{false}, smi: map[string]interface{}{"false": false}},
		{input: int64(3306), s: "3306", b: true, i: int64(3306), is: []interface{}{"3306"}, smi: map[string]interface{}{"0": "3306"}},
		{input: []string{"Hello", "World", "!"}, s: "[\"Hello\",\"World\",\"!\"]", b: true, i: int64(1), is: []interface{}{"Hello", "World", "!"}, smi: map[string]interface{}{"0": "Hello", "1": "World", "2": "!"}},
		{input: []int{1, 2, 3}, s: "[1,2,3]", b: true, i: int64(1), is: []interface{}{1, 2, 3}, smi: map[string]interface{}{"0": 1, "1": 2, "2": 3}},
		{input: []interface{}{"Hello", "World", 123}, s: "[\"Hello\",\"World\",123]", b: true, i: int64(1), is: []interface{}{"Hello", "World", 123}, smi: map[string]interface{}{"0": "Hello", "1": "World", "2": 123}},
		{input: map[string]string{"hi": "HelloWorld", "name": "Kisschou"}, s: "{\"hi\":\"HelloWorld\",\"name\":\"Kisschou\"}", b: true, i: int64(1), is: []interface{}{"hi", "HelloWorld", "name", "Kisschou"}, smi: map[string]interface{}{"hi": "HelloWorld", "name": "Kisschou"}},
		{input: map[string]int{"name": 1, "age": 13}, s: "{\"age\":13,\"name\":1}", b: true, i: int64(1), is: []interface{}{"name", 1, "age", 13}, smi: map[string]interface{}{"name": 1, "age": 13}},
		{input: map[string]interface{}{"name": 1, "age": 13, "hi": "HelloWorld"}, s: "{\"age\":13,\"hi\":\"HelloWorld\",\"name\":1}", b: true, i: int64(1), is: []interface{}{"name", 1, "age", 13, "hi", "HelloWorld"}, smi: map[string]interface{}{"name": 1, "age": 13, "hi": "HelloWorld"}},
		{input: map[string][]string{"s": []string{"Hi", "hi", "hI"}, "v": []string{"v1.0"}}, s: "{\"s\":[\"Hi\",\"hi\",\"hI\"],\"v\":[\"v1.0\"]}", b: true, i: int64(1), is: []interface{}{"s", []string{"Hi", "hi", "hI"}, "v", []string{"v1.0"}}, smi: map[string]interface{}{"s": []string{"Hi", "hi", "hI"}, "v": []string{"v1.0"}}},
	}

	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			tStruct := New(test.input)
			if tStruct.String != test.s {
				fmt.Println(fmt.Sprintf("Fail verify string. got: %v, want: %v.", tStruct.String, test.s))
			}
			if tStruct.Bool != test.b {
				fmt.Println(fmt.Sprintf("Fail verify bool Fail. got: %v, want: %v.", tStruct.Bool, test.b))
			}
			if tStruct.Int64 != test.i {
				fmt.Println(fmt.Sprintf("Fail verify int64 Fail. got: %v, want: %v.", tStruct.Int64, test.i))
			}
			if !SliceEqual(tStruct.InterfaceSlice, test.is) {
				fmt.Println(fmt.Sprintf("Fail verify slice. got: %v, want: %v.", tStruct.InterfaceSlice, test.is))
			}
			if !MapEqual(tStruct.StringMapInterface, test.smi) {
				fmt.Println(fmt.Sprintf("Fail verify map. got: %v, want: %v.", tStruct.StringMapInterface, test.smi))
			}
		}
	}
}

// <=========
