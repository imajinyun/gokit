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

func TestEmpty(t *testing.T) {
	var emptyTests = []struct {
		in  any
		out bool
	}{
		{"", true},
		{"0", false},
    {".0", false},
    {false, true},
    {true, false},
		{nil, true},
		{0, true},
    {0.1, false},
    {3.14, false},
    {uint(0), true},
    {int64(-1), false},
    {[]string{}, true},
    {[]int{}, true},
    {[...]string{}, true},
    {[...]int{}, true},
    {map[string]any{}, true},
    {make(map[string]int), true},
	}

	for _, tt := range emptyTests {
		if got := Empty(tt.in); got != tt.out {
			t.Errorf("Empty(%v) = %v, want: %v", tt.in, got, tt.out)
		}
	}
}
