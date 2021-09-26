package main

import "testing"

func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func TestRotate(t *testing.T) {
	var tests = []struct {
		input  []int
		count  int
		result []int
	}{
		{[]int{}, 2, []int{}},
		{[]int{1}, 3, []int{1}},
		{[]int{0, 1}, 3, []int{1, 0}},
		{[]int{0, 1, 2, 3, 4, 5}, 0, []int{0, 1, 2, 3, 4, 5}},
		{[]int{0, 1, 2, 3, 4, 5}, 2, []int{2, 3, 4, 5, 0, 1}},
		{[]int{0, 1, 2, 3, 4, 5, 6}, 10, []int{3, 4, 5, 6, 0, 1, 2}},
	}

	for _, test := range tests {
		tmp := make([]int, len(test.input))
		copy(tmp, test.input)
		rotate(test.input, test.count)
		if !equal(test.input, test.result) {
			t.Errorf("rotate(%v, %v), then %v\n", tmp, test.count, test.input)
		}
	}

}
