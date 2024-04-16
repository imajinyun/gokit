package gohelper

import (
	"testing"
)

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
