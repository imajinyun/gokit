package gohelper

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.Config{SortMapKeys: true}.Froze()

func TestToString(t *testing.T) {
	tests := []struct {
		data any
		out  string
	}{
		{"hello world", "hello world"},
		{true, "true"},
    {false, "false"},
		{0, "0"},
		{123, "123"},
		{3.1415926, "3.1415926"},
		{-3.1415926, "-3.1415926"},
		{nil, "<nil>"},
		{"", ""},
		{[]int{1, 2, 3}, "[1 2 3]"},
		{[]float64{1.20, 2.30, 3.40}, "[1.2 2.3 3.4]"},
		{[...]uint{1, 2, 3}, "[1 2 3]"},
		{struct{ name string }{"jack"}, `{jack}`},
		{map[string]any{"hello": "world", "world": 123}, `map[hello:world world:123]`},
	}

	for _, tt := range tests {
		got := ToString(tt.data)
		if got != tt.out {
			t.Errorf("ToString(%v) = %v, want: %v", tt.data, got, tt.out)
		}
	}
}

func TestStructToMap(t *testing.T) {
	tests := []struct {
		name string
		data any
		want string
	}{
		{
			name: "test case 1",
			data: struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			}{1, "jack"},
			want: `{"id":1,"name":"jack"}`,
		},
		{
			name: "test case 2",
			data: struct {
				TestID   int    `json:"test_id"`
				TestName string `json:"test_name"`
			}{10, "test name"},
			want: `{"test_id":10,"test_name":"test name"}`,
		},
		{
			name: "test case 3",
			data: struct{}{},
			want: "{}",
		},
		{
			name: "test case 4",
			data: struct {
				Name      string `json:"name"`
				Age       uint8  `json:"age"`
				Email     string `json:"email"`
				IsMarried bool   `json:"is_married"`
				CreatedAt string `json:"created_at"`
			}{"jack", 18, "jack@example.com", false, "2024-03-13 18:00:00"},
			want: `{"age":18,"created_at":"2024-03-13 18:00:00","email":"jack@example.com","is_married":false,"name":"jack"}`,
		},
		{
			name: "test case 5",
			data: struct {
				StrField  string `json:"str_field"`
				NumField  int    `json:"num_field"`
				BoolField bool   `json:"bool_field"`
				NestField struct {
					F1 string `json:"f1"`
					F2 int    `json:"f2"`
				} `json:"nest_field"`
			}{},
			want: `{"bool_field":false,"nest_field":{"f1":"","f2":0},"num_field":0,"str_field":""}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("TestStructToMap() panicked: %v", r)
				}
			}()

			got := StructToMap[any](tt.data)
			byt, err := json.Marshal(got)
			if err != nil {
				t.Errorf("TestStructToMap() error: %v", err)
			}

			if string(byt) != tt.want {
				t.Errorf("StructToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
