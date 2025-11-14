package jsonliteparser

import (
	"reflect"
	"testing"
)

func TestEmptyStructures(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  interface{}
	}{
		{"empty object", `{}`, map[string]interface{}{}},
		{"empty array", `[]`, []interface{}{}},
		{"empty object with spaces", `{  }`, map[string]interface{}{}},
		{"empty array with spaces", `[  ]`, []interface{}{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJSON(tt.input)
			if err != nil {
				t.Errorf("ParseJSON(%q) error = %v", tt.input, err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseJSON(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestSimpleObject(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  map[string]interface{}
	}{
		{
			"single string property",
			`{"name":"John"}`,
			map[string]interface{}{"name": "John"},
		},
		{
			"single number property",
			`{"age":30}`,
			map[string]interface{}{"age": 30.0},
		},
		{
			"single boolean property",
			`{"active":true}`,
			map[string]interface{}{"active": true},
		},
		{
			"null property",
			`{"value":null}`,
			map[string]interface{}{"value": nil},
		},
		{
			"multiple properties",
			`{"name":"John","age":30,"active":true}`,
			map[string]interface{}{"name": "John", "age": 30.0, "active": true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJSON(tt.input)
			if err != nil {
				t.Errorf("ParseJSON(%q) error = %v", tt.input, err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseJSON(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestNestedObject(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  interface{}
	}{
		{
			"one level nesting",
			`{"user":{"name":"John"}}`,
			map[string]interface{}{
				"user": map[string]interface{}{"name": "John"},
			},
		},
		{
			"two level nesting",
			`{"data":{"user":{"name":"John"}}}`,
			map[string]interface{}{
				"data": map[string]interface{}{
					"user": map[string]interface{}{"name": "John"},
				},
			},
		},
		{
			"multiple nested objects",
			`{"user":{"name":"John","address":{"city":"NYC"}}}`,
			map[string]interface{}{
				"user": map[string]interface{}{
					"name": "John",
					"address": map[string]interface{}{
						"city": "NYC",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJSON(tt.input)
			if err != nil {
				t.Errorf("ParseJSON(%q) error = %v", tt.input, err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseJSON(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestSimpleArray(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []interface{}
	}{
		{
			"number array",
			`[1,2,3]`,
			[]interface{}{1.0, 2.0, 3.0},
		},
		{
			"string array",
			`["a","b","c"]`,
			[]interface{}{"a", "b", "c"},
		},
		{
			"boolean array",
			`[true,false,true]`,
			[]interface{}{true, false, true},
		},
		{
			"mixed array",
			`[1,"text",true,null]`,
			[]interface{}{1.0, "text", true, nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJSON(tt.input)
			if err != nil {
				t.Errorf("ParseJSON(%q) error = %v", tt.input, err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseJSON(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestArrayOfObjects(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  interface{}
	}{
		{
			"array of simple objects",
			`[{"name":"John"},{"name":"Jane"}]`,
			[]interface{}{
				map[string]interface{}{"name": "John"},
				map[string]interface{}{"name": "Jane"},
			},
		},
		{
			"array of complex objects",
			`[{"name":"John","age":30},{"name":"Jane","age":25}]`,
			[]interface{}{
				map[string]interface{}{"name": "John", "age": 30.0},
				map[string]interface{}{"name": "Jane", "age": 25.0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJSON(tt.input)
			if err != nil {
				t.Errorf("ParseJSON(%q) error = %v", tt.input, err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseJSON(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestNestedArrays(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  interface{}
	}{
		{
			"array in array",
			`[[1,2],[3,4]]`,
			[]interface{}{
				[]interface{}{1.0, 2.0},
				[]interface{}{3.0, 4.0},
			},
		},
		{
			"object with array",
			`{"numbers":[1,2,3]}`,
			map[string]interface{}{
				"numbers": []interface{}{1.0, 2.0, 3.0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJSON(tt.input)
			if err != nil {
				t.Errorf("ParseJSON(%q) error = %v", tt.input, err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseJSON(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestWhitespace(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  interface{}
	}{
		{
			"spaces around values",
			`{ "name" : "John" }`,
			map[string]interface{}{"name": "John"},
		},
		{
			"newlines",
			"{\n\"name\":\"John\"\n}",
			map[string]interface{}{"name": "John"},
		},
		{
			"tabs",
			"{\t\"name\":\t\"John\"\t}",
			map[string]interface{}{"name": "John"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJSON(tt.input)
			if err != nil {
				t.Errorf("ParseJSON(%q) error = %v", tt.input, err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseJSON(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestNumbers(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  interface{}
	}{
		{"positive integer", `{"n":123}`, map[string]interface{}{"n": 123.0}},
		{"negative integer", `{"n":-45}`, map[string]interface{}{"n": -45.0}},
		{"decimal", `{"n":3.14}`, map[string]interface{}{"n": 3.14}},
		{"negative decimal", `{"n":-2.5}`, map[string]interface{}{"n": -2.5}},
		{"zero", `{"n":0}`, map[string]interface{}{"n": 0.0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJSON(tt.input)
			if err != nil {
				t.Errorf("ParseJSON(%q) error = %v", tt.input, err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseJSON(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestErrorCases(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"empty string", ``},
		{"unclosed object", `{"name":"John"`},
		{"unclosed array", `[1,2,3`},
		{"unclosed string", `{"name":"John}`},
		{"missing colon", `{"name""John"}`},
		{"missing comma in object", `{"a":1"b":2}`},
		{"missing comma in array", `[1 2]`},
		{"trailing comma in object", `{"name":"John",}`},
		{"trailing comma in array", `[1,2,3,]`},
		{"invalid value", `{"name":undefined}`},
		{"single quote string", `{'name':'John'}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseJSON(tt.input)
			if err == nil {
				t.Errorf("ParseJSON(%q) expected error, got nil", tt.input)
			}
		})
	}
}

func TestComplexStructure(t *testing.T) {
	input := `{
		"name": "John",
		"age": 30,
		"address": {
			"street": "123 Main St",
			"city": "NYC"
		},
		"hobbies": ["reading", "coding"],
		"active": true,
		"metadata": null
	}`

	want := map[string]interface{}{
		"name": "John",
		"age":  30.0,
		"address": map[string]interface{}{
			"street": "123 Main St",
			"city":   "NYC",
		},
		"hobbies":  []interface{}{"reading", "coding"},
		"active":   true,
		"metadata": nil,
	}

	got, err := ParseJSON(input)
	if err != nil {
		t.Errorf("ParseJSON() error = %v", err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ParseJSON() = %v, want %v", got, want)
	}
}
