package gohelper

import "testing"

func TestRandNum(t *testing.T) {
	var randNumTests = []struct {
		min int
		max int
	}{
		{1, 9},
		{10, 99},
		{100, 999},
		{1000, 9999},
		{10000, 99999},
		{100000, 999999},
		{1000000, 9999999},
		{10000000, 99999999},
	}

	for _, tt := range randNumTests {
		got := RandInt(tt.min, tt.max)
		if got < tt.min || got >= tt.max {
			t.Errorf("Number %d is out of range [%d,%d)", got, tt.min, tt.max)
		}
	}
}
