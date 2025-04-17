package maps

import "testing"

func TestGetOrDefault(t *testing.T) {
	var getOrDefaultTests = []struct {
		m   map[any]any
		key any
		def any
		out any
	}{
		{map[any]any{"a": 1, "b": 2, "c": 3}, "a", 1, 1},
		{map[any]any{"a": 1, "b": 2, "c": 3}, "d", 4, 4},
		{map[any]any{"1": "a", "2": "b", "3": "c"}, "1", "a", "a"},
		{map[any]any{"1": "a", "2": "b", "3": "c"}, "4", "4", "4"},
		{map[any]any{"a": true, "b": false, "c": true}, "a", true, true},
		{map[any]any{"a": true, "b": false, "c": true}, "d", false, false},
		{map[any]any{1: 11, 2: 22, 3: 33}, 1, 11, 11},
		{map[any]any{1: 11, 2: 22, 3: 33}, 1, int(11), int(11)},
		{map[any]any{1: uint(11), 2: uint(22), 3: uint(33)}, 4, uint(33), uint(33)},
		{map[any]any{1: uint(11), 2: uint(22), 3: uint(33)}, 4, uint(44), uint(44)},
		{map[any]any{1: int8(11), 2: int8(22), 3: int8(33)}, 1, int8(11), int8(11)},
		{map[any]any{1: uint8(11), 2: uint8(22), 3: uint8(33)}, 4, uint8(44), uint8(44)},
		{map[any]any{1: int16(11), 2: int16(22), 3: int16(33)}, 1, int16(11), int16(11)},
		{map[any]any{1: uint16(11), 2: uint16(22), 3: uint16(33)}, 4, uint16(44), uint16(44)},
		{map[any]any{1: int32(11), 2: int32(22), 3: int32(33)}, 1, int32(11), int32(11)},
		{map[any]any{1: uint32(11), 2: uint32(22), 3: uint32(33)}, 4, uint32(44), uint32(44)},
		{map[any]any{1: int64(11), 2: int64(22), 3: int64(33)}, 1, int64(11), int64(11)},
		{map[any]any{1: uint64(11), 2: uint64(22), 3: uint64(33)}, 4, uint64(44), uint64(44)},
		{map[any]any{uint(1): 11, uint(2): 22, uint(3): 33}, uint(1), 11, 11},
		{map[any]any{uint(1): 11, uint(2): 22, uint(3): 33}, uint(1), int(11), int(11)},
		{map[any]any{uint(1): uint(11), uint(2): uint(22), uint(3): uint(33)}, uint(4), uint(44), uint(44)},
		{map[any]any{uint(1): int8(11), uint(2): int8(22), uint(3): int8(33)}, uint(1), int8(11), int8(11)},
		{map[any]any{uint(1): uint8(11), uint(2): uint8(22), uint(3): uint8(33)}, uint(4), uint8(44), uint8(44)},
		{map[any]any{uint(1): int16(11), uint(2): int16(22), uint(3): int16(33)}, uint(1), int16(11), int16(11)},
		{map[any]any{uint(1): uint16(11), uint(2): uint16(22), uint(3): uint16(33)}, uint(4), uint16(44), uint16(44)},
		{map[any]any{uint(1): int32(11), uint(2): int32(22), uint(3): int32(33)}, uint(1), int32(11), int32(11)},
		{map[any]any{uint(1): uint32(11), uint(2): uint32(22), uint(3): uint32(33)}, uint(4), uint32(44), uint32(44)},
		{map[any]any{uint(1): int64(11), uint(2): int64(22), uint(3): int64(33)}, uint(1), int64(11), int64(11)},
		{map[any]any{uint(1): uint64(11), uint(2): uint64(22), uint(3): uint64(33)}, uint(4), uint64(44), uint64(44)},
	}
	for _, tt := range getOrDefaultTests {
		got := GetOrDefault(tt.m, tt.key, tt.def)
		if got != tt.out {
			t.Errorf("GetOrDefault(%v, %v, %v) = %v, want %v", tt.m, tt.key, tt.def, got, tt.out)
		}
	}
}
