package gohelper

import (
	"fmt"
	"testing"
)

func TestIf(t *testing.T) {
	var ifTests = []struct {
		expr bool
		succ any
		fail any
		out  any
	}{
		{1 == 1, true, false, true},
		{2 == 2, true, false, true},
		{true, "ok", "no", "ok"},
		{0 > -1, 123, 456, 123},
		{1 != 1, "equal", "not equal", "not equal"},
		{3.14 == 3.14, true, false, true},
		{-3.14 != -3.14, true, false, false},
		{"🐶" == "🐶", true, false, true},
	}

	for _, tt := range ifTests {
		got := If(tt.expr, tt.succ, tt.fail)
		switch got.(type) {
		case bool:
			if got.(bool) != tt.out {
				t.Errorf("If(%v, %v, %v) = %v, want: %v", tt.expr, tt.succ, tt.fail, got.(bool), tt.out)
			}
		case string:
			if got.(string) != tt.out {
				t.Errorf("If(%v, %v, %v) = %v, want: %v", tt.expr, tt.succ, tt.fail, got.(string), tt.out)
			}
		case int, int8, int16, int32, int64:
			if fmt.Sprintf("%d", got) != fmt.Sprintf("%d", tt.out) {
				t.Errorf("If(%v, %v, %v) = %v, want: %v", tt.expr, tt.succ, tt.fail, got, tt.out)
			}
		case uint, uint8, uint16, uint32, uint64:
			if fmt.Sprintf("%d", got) != fmt.Sprintf("%d", tt.out) {
				t.Errorf("If(%v, %v, %v) = %v, want: %v", tt.expr, tt.succ, tt.fail, got, tt.out)
			}
		default:
			t.Error("unsupported type")
		}
	}
}
