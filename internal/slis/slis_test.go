package slis

import "testing"

func TestFilter(t *testing.T) {
	type test[T any] struct {
		name string
		data []T
		iter func(i int, v T) bool
		out  []T
	}
	tests := []test[any]{
		{
			"test case 1",
			[]any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			func(i int, v any) bool { return v.(int)%2 == 0 },
			[]any{2, 4, 6, 8, 10},
		},
		{
			"test case 2",
			[]any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			func(i int, v any) bool { return v.(int)%2 == 1 },
			[]any{1, 3, 5, 7, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.data, tt.iter)
			for i := range got {
				if got[i] != tt.out[i] {
					t.Errorf("Filter(%v) = %v, want: %v", tt.data, got, tt.out)
				}
			}
		})
	}

	{
		type demo[T uint] struct {
			name string
			data []T
			iter func(i int, v T) bool
			out  []T
		}
		tt := demo[uint]{
			"test case 3",
			[]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			func(i int, v uint) bool { return v%2 == 0 },
			[]uint{2, 4, 6, 8, 10},
		}

		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.data, tt.iter)
			for i := range got {
				if got[i] != tt.out[i] {
					t.Errorf("Filter(%v) = %v, want: %v", tt.data, got, tt.out)
				}
			}
		})
	}

	{
		type demo[T string] struct {
			name string
			data []T
			iter func(i int, v T) bool
			out  []T
		}
		tt := demo[string]{
			"test case 4",
			[]string{"hello", "world", "test", "demo", "golang"},
			func(i int, v string) bool { return len(v) == 4 },
			[]string{"test", "demo"},
		}

		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.data, tt.iter)
			for i := range got {
				if got[i] != tt.out[i] {
					t.Errorf("Filter(%v) = %v, want: %v", tt.data, got, tt.out)
				}
			}
		})
	}
}

func TestInSlice(t *testing.T) {
	var tests = []struct {
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

	for _, tt := range tests {
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
