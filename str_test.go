package gohelper

import "testing"

func TestRandStr(t *testing.T) {
	var randStrTests = []struct {
		len int
	}{
		{1},
		{10},
		{100},
		{1000},
		{10000},
		{100000},
		{1000000},
		{10000000},
	}

	for _, tt := range randStrTests {
		got := RandStr(tt.len)
		if len(got) != tt.len {
			t.Errorf("RandStr(%d) = %d, want: %d", tt.len, len(got), tt.len)
		}
	}
}
