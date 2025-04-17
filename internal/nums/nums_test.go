package nums

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
		t.Log(got)
		if got < tt.min || got >= tt.max {
			t.Errorf("Number %d is out of range [%d,%d)", got, tt.min, tt.max)
		}
	}
}

func TestIsDigit(t *testing.T) {
	var tests = []struct {
		str string
		out bool
	}{
		{"123", true},
		{"-123", false},
		{"0.123", false},
		{"-0.123", false},
		{"0", true},
		{"0.00", false},
		{".123", false},
		{"123.", false},
		{"123abc", false},
		{"a1b2", false},
	}

	for _, tt := range tests {
		got := IsDigit(tt.str)
		if got != tt.out {
			t.Errorf("IsNumber(%s) = %v, want: %v", tt.str, got, tt.out)
		}
	}
}

func TestIsNumber(t *testing.T) {
	var tests = []struct {
		str string
		out bool
	}{
		{"123", true},
		{"-123", true},
		{"0.123", true},
		{"-0.123", true},
		{"0", true},
		{"0.00", true},
		{".0", false},
		{".123", false},
		{"123abc", false},
		{"a1b2", false},
	}

	for _, tt := range tests {
		got := IsNumber(tt.str)
		if got != tt.out {
			t.Errorf("IsNumber(%s) = %v, want: %v", tt.str, got, tt.out)
		}
	}
}

func TestIsFloat(t *testing.T) {
	var tests = []struct {
		str string
		out bool
	}{
		{"123", true},
		{"-123", true},
		{"0.123", true},
		{"-0.123", true},
		{"0", true},
		{"0.00", true},
		{".0", true},
		{".123", true},
		{"123abc", false},
		{"a1b2", false},
	}

	for _, tt := range tests {
		got := IsFloat(tt.str)
		if got != tt.out {
			t.Errorf("IsFloat(%s) = %v, want: %v", tt.str, got, tt.out)
		}
	}
}

func TestIsInteger(t *testing.T) {
	var tests = []struct {
		num any
		out bool
	}{
		{123, true},
		{-123, true},
		{0.123, false},
		{-0.123, false},
		{0, true},
		{"0", true},
		{0.00, true},
		{"0.00", true},
    {.0, true},
    {".0", true},
    {true, true},
    {false, true},
	}

	for _, tt := range tests {
		got := IsInteger(tt.num)
		if got != tt.out {
			t.Errorf("IsInteger(%s) = %v, want: %v", tt.num, got, tt.out)
		}
	}
}

func TestIsNatural(t *testing.T) {
	var tests = []struct {
		val any
		out bool
	}{
		{123, true},
		{"123", true},
		{123.123, false},
		{-123, false},
		{"-123", true},
		{-123.123, false},
		{0.123, false},
		{"0.123", false},
		{-0.123, false},
		{"-0.123", false},
		{0, true},
		{"0", true},
		{false, false},
		{true, true},
	}

	for _, tt := range tests {
		got := IsNatural(tt.val)
		if got != tt.out {
			t.Errorf("IsNatural(%s) = %v, want: %v", tt.val, got, tt.out)
		}
	}
}
