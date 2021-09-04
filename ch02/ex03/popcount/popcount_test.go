package popcount

import "testing"

func TestPopCount(t *testing.T) {
	var tests = []struct {
		input uint64
		want  int
	}{
		{0, 0},
		{18446744073709551615, 64},
		{13, 3},
	}
	for _, test := range tests {
		if got := PopCount(test.input); got != test.want {
			t.Errorf("PopCount(%d) = %d\n", test.input, got)
		}
	}
}
