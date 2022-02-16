package equal

import (
	"testing"
)

func TestEqual(t *testing.T) {
	tests := []struct {
		x, y interface{}
		want bool
	}{
		{int8(0), int8(0), true},
		{int8(0), int16(0), false},
		{uint32(1), uint32(1), true},
		{1, 2, false},
		{int16(10), int16(2), false},
		{float64(0.1), float64(0.2), false},
		{float32(0.1), float32(0.1), true},
		{0.0000000011, 0.000000001, true},
		{0.0030000054, 0.0030000058, true},
		{0.003000054, 0.003000058, false},
		{complex(1, 2), complex(1, 2), true},
		{complex(1, 3.0002060056), complex(1, 3.0002060062), true},
		{complex(1, 3.0002060056), complex(1, 3.0002060067), false},
	}

	for _, test := range tests {
		if got := Equal(test.x, test.y); got != test.want {
			t.Errorf("not equal: %v, %v\n", test.x, test.y)
		}
	}
}
