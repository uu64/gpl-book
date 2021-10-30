package main

import "testing"

func TestMinWithError(t *testing.T) {
	tests := []struct {
		input   []int
		want    int
		isError bool
	}{
		{[]int{}, 0, true},
		{[]int{2}, 2, false},
		{[]int{4, 2, 3}, 2, false},
	}
	for _, test := range tests {
		got, err := minWithError(test.input...)
		if test.isError {
			if err == nil {
				t.Errorf("must return error\n")
			}
			continue
		} else {
			if err != nil {
				t.Errorf("error: minWithError(%v)\n", test.input)
			}
			if got != test.want {
				t.Errorf("error: minWithError(%v) returns %v\n", test.input, got)
			}
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		input1 int
		input2 []int
		want   int
	}{
		{2, []int{}, 2},
		{5, []int{4}, 4},
		{1, []int{4, 2, 3}, 1},
	}
	for _, test := range tests {
		got := min(test.input1, test.input2...)
		if got != test.want {
			t.Errorf("error: min(%v, %v) returns %v\n", test.input1, test.input2, got)
		}
	}
}

func TestMaxWithError(t *testing.T) {
	tests := []struct {
		input   []int
		want    int
		isError bool
	}{
		{[]int{}, 0, true},
		{[]int{2}, 2, false},
		{[]int{4, 2, 3}, 4, false},
	}
	for _, test := range tests {
		got, err := maxWithError(test.input...)
		if test.isError {
			if err == nil {
				t.Errorf("must return error\n")
			}
			continue
		} else {
			if err != nil {
				t.Errorf("error: maxWithError(%v)\n", test.input)
			}
			if got != test.want {
				t.Errorf("error: maxWithError(%v) returns %v\n", test.input, got)
			}
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		input1 int
		input2 []int
		want   int
	}{
		{2, []int{}, 2},
		{5, []int{4}, 5},
		{1, []int{4, 2, 3}, 4},
	}
	for _, test := range tests {
		got := max(test.input1, test.input2...)
		if got != test.want {
			t.Errorf("error: max(%v, %v) returns %v\n", test.input1, test.input2, got)
		}
	}
}
