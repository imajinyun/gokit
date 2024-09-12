package gohelper

import (
	"testing"
)

func TestRandStr(t *testing.T) {
	var tests = []struct {
		len int
		out int
	}{
		{1, 1},
		{2, 2},
		{4, 4},
		{8, 8},
		{16, 16},
		{32, 32},
		{64, 64},
		{128, 128},
		{256, 256},
		{512, 512},
		{1024, 1024},
		{2048, 2048},
	}

	for _, tt := range tests {
		got := RandStr(tt.len)
		if len(got) != tt.out {
			t.Errorf("RandStr(%d) = %d, want: %d", tt.len, len(got), tt.out)
		}
	}
}

func TestRandStrWithOption(t *testing.T) {
	var tests = []struct {
		len int
		out int
	}{
		{1, 1},
		{2, 2},
		{4, 4},
		{8, 8},
		{16, 16},
		{32, 32},
		{64, 64},
		{128, 128},
		{256, 256},
		{512, 512},
		{1024, 1024},
		{2048, 2048},
	}

	for _, tt := range tests {
		got := RandStrWithOption(tt.len, Option{
			IncludeNumber:    true,
			IncludeUppercase: true,
			IncludeSymbol:    true,
		})
		if len(got) != tt.out {
			t.Errorf("RandStr(%d) = %d, want: %d", tt.len, len(got), tt.out)
		}
	}
}

func TestEmpty(t *testing.T) {
	var tests = []struct {
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

	for _, tt := range tests {
		if got := Empty(tt.in); got != tt.out {
			t.Errorf("Empty(%v) = %v, want: %v", tt.in, got, tt.out)
		}
	}
}

func TestTrim(t *testing.T) {
	tests := []struct {
		name string
		s    string
		char string
		out  string
	}{
		{name: "test case 1", s: "hello", char: "", out: "hello"},
		{name: "trim case 2", s: "  hello    ", char: " ", out: "hello"},
		{name: "trim case 3", s: "  hello    ", char: "", out: "hello"},
		{name: "trim case 4", s: "  hello    ", char: "h", out: "  hello    "},
		{name: "trim case 5", s: "  hello    ", char: "h ", out: "ello"},
		{name: "trim case 6", s: "  hello    ", char: " h", out: "ello"},
		{name: "trim case 7", s: "  hello    ", char: "lo", out: "  hello    "},
		{name: "trim case 8", s: "  hello    ", char: "lo ", out: "he"},
		{name: "trim case 9", s: "  hello    ", char: " lo", out: "he"},
		{name: "trim case 10", s: "hello ", char: "he", out: "llo "},
		{name: "trim case 11", s: "12345678987654321", char: "1-8", out: "9"},
		{name: "trim case 12", s: "040400204000", char: "04", out: "2"},
		{name: "trim case 13", s: "  \r\nhello \r\n \n \t  ", char: "", out: "hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Trim(tt.s, tt.char)
			if got != tt.out {
				t.Errorf("Trim(%v, %v) = %v, want: %v", tt.s, tt.char, got, tt.out)
			}
		})
	}
}

func TestTrimLeft(t *testing.T) {
	tests := []struct {
		name string
		s    string
		char string
		out  string
	}{
		{name: "test case 1", s: "hello", char: "", out: "hello"},
		{name: "trim case 2", s: "  hello    ", char: " ", out: "hello    "},
		{name: "trim case 3", s: "  hello    ", char: "", out: "hello    "},
		{name: "trim case 4", s: "  hello    ", char: "h", out: "  hello    "},
		{name: "trim case 5", s: "  hello    ", char: "h ", out: "ello    "},
		{name: "trim case 6", s: "  hello    ", char: " h", out: "ello    "},
		{name: "trim case 7", s: "  hello    ", char: "lo", out: "  hello    "},
		{name: "trim case 8", s: "  hello    ", char: "lo ", out: "hello    "},
		{name: "trim case 9", s: "  hello    ", char: " lo", out: "hello    "},
		{name: "trim case 10", s: "hello ", char: "he", out: "llo "},
		{name: "trim case 11", s: "12345678987654321", char: "1-8", out: "987654321"},
		{name: "trim case 12", s: "040400204000", char: "04", out: "204000"},
		{name: "trim case 13", s: "  \r\nhello \r\n \n \t  ", char: "", out: "hello \r\n \n \t  "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TrimLeft(tt.s, tt.char)
			if got != tt.out {
				t.Errorf("TrimLeft(%v, %v) = %v, want: %v", tt.s, tt.char, got, tt.out)
			}
		})
	}
}

func TestTrimRight(t *testing.T) {
	tests := []struct {
		name string
		s    string
		char string
		out  string
	}{
		{name: "test case 1", s: "hello", char: "", out: "hello"},
		{name: "trim case 2", s: "  hello    ", char: " ", out: "  hello"},
		{name: "trim case 3", s: "  hello    ", char: "", out: "  hello"},
		{name: "trim case 4", s: "  hello    ", char: "h", out: "  hello    "},
		{name: "trim case 5", s: "  hello    ", char: "h ", out: "  hello"},
		{name: "trim case 6", s: "  hello    ", char: " h", out: "  hello"},
		{name: "trim case 7", s: "  hello    ", char: "lo", out: "  hello    "},
		{name: "trim case 8", s: "  hello    ", char: "lo ", out: "  he"},
		{name: "trim case 9", s: "  hello    ", char: " lo", out: "  he"},
		{name: "trim case 10", s: "hello ", char: "he", out: "hello "},
		{name: "trim case 11", s: "12345678987654321", char: "1-8", out: "123456789"},
		{name: "trim case 12", s: "040400204000", char: "04", out: "0404002"},
		{name: "trim case 13", s: "  \r\nhello \r\n \n \t  ", char: "", out: "  \r\nhello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TrimRight(tt.s, tt.char)
			if got != tt.out {
				t.Errorf("TrimRight(%v, %v) = %v, want: %v", tt.s, tt.char, got, tt.out)
			}
		})
	}
}
