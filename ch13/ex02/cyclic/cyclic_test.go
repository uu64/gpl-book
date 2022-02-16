package cyclic

import (
	"fmt"
	"testing"
)

type link struct {
	value string
	tail  *link
}

func TestIsCyclic(t *testing.T) {
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c

	tests := []struct {
		x    interface{}
		want bool
	}{
		{1, false},
		{2.345, false},
		{complex(2, 3), false},
		{"test", false},
		{true, false},
		{nil, false},
		{a, true},
		{c, true},
	}

	for _, test := range tests {
		if got := IsCyclic(test.x); got != test.want {
			t.Errorf("IsCyclic(%v) returns %t, want %t\n", test.x, got, test.want)
		} else {
			fmt.Printf("IsCyclic(%v) returns %t, want %t\n", test.x, got, test.want)
		}
	}
}
