package gohelper

import "testing"

func TestInSlice(t *testing.T) {
	var inSliceTests = []struct {
		in    any
		slice []any
		out   any
	}{
		{"a", []any{"a", "b", "c"}, true},
		{"d", []any{"a", "b", "c"}, false},
		{1, []any{1, 2, 3}, true},
		{4, []any{1, 2, 3}, false},
		{"hello", []any{"a", 4, "hello", true, false}, true},
		{true, []any{false, false, false}, false},
		{false, []any{false, false, false}, true},
		{1.01, []any{1.01, 1.99, 3.14}, true},
		{
			"/api/v1/upload/image",
			[]any{
				"/api/v1/upload/image",
				"/api/v2/upload/image",
				"/api/v3/upload/image",
			},
			true,
		},
	}

	for _, tt := range inSliceTests {
		got := InSlice(tt.in, tt.slice)
		if got != tt.out {
			t.Errorf("InSlice(%v, %v) = %v, want %v", tt.in, tt.slice, got, tt.out)
		}
	}

	t.Run("ElementPresent", func(t *testing.T) {
		list := []int{1, 2, 3, 4, 5}
		elem := 3
		if !InSlice(elem, list) {
			t.Errorf("Expected true, got false")
		}
	})

	t.Run("ElementNotPresent", func(t *testing.T) {
		list := []int{1, 2, 3, 4, 5}
		elem := 6
		if InSlice(elem, list) {
			t.Errorf("Expected false, got true")
		}
	})

	t.Run("EmptySlice", func(t *testing.T) {
		list := []int{}
		elem := 5
		if InSlice(elem, list) {
			t.Errorf("Expected false, got true")
		}
	})
}
