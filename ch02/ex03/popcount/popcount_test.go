package popcount

import "testing"

func TestPopCount(t *testing.T) {
	var tests = []struct {
		input uint64
		want  int
	}{
		{0, 0},
		{255, 8},
		{13, 3},
	}
	for _, test := range tests {
		if got := PopCount(test.input); got != test.want {
			t.Errorf("PopCount(%d) = %d\n", test.input, got)
		}
	}
}
