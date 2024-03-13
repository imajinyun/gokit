package gohelper

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.Config{SortMapKeys: true}.Froze()

func TestStructToMap(t *testing.T) {
	var mapTests = []struct {
		Name      string `json:"name"`
		Age       uint8  `json:"age"`
		Email     string `json:"email"`
		IsMarried bool   `json:"is_married"`
		CreatedAt string `json:"created_at"`
	}{
		{"jack", 18, "jack@example.com", false, "2024-03-13 18:00:00"},
		{"pony", 20, "pony@example.com", true, "2024-03-13 18:31:23"},
	}

	outs := []string{
		`{"age":18,"created_at":"2024-03-13 18:00:00","email":"jack@example.com","is_married":false,"name":"jack"}`,
		`{"age":20,"created_at":"2024-03-13 18:31:23","email":"pony@example.com","is_married":true,"name":"pony"}`,
	}

	for i, tt := range mapTests {
		got := StructToMap(tt)
		byt, err := json.MarshalToString(got)
		if err != nil {
			t.Fatal(err)
		}

		out := outs[i]
		if string(byt) != out {
			t.Errorf("StructToMap(%v) = %v, want: %v", tt, string(byt), out)
		}
	}
}
