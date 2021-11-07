package treesort

import "testing"

func TestString(t *testing.T) {
	var root *tree
	if got := root.String(); got != "[]" {
		t.Errorf("want: %v, got: %v\n", "", got)
	}

	list := []int{2, 15, 7, 3, 9, 10}
	for _, v := range list {
		root = add(root, v)
	}
	if got := root.String(); got != "[2 3 7 9 10 15]" {
		t.Errorf("want: %v, got: %v\n", "", got)
	}
}
