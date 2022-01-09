package intset

import (
	"fmt"
	"testing"
)

func TestIntSet_one(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	if s := x.String(); s != "{1 9 144}" {
		t.Errorf("x.String() got: %s, expected: {1 9 144}\n", s)
	}

	y.Add(9)
	y.Add(42)
	if s := y.String(); s != "{9 42}" {
		t.Errorf("y.String() got: %s, expected: {9 42}\n", s)
	}

	x.UnionWith(&y)
	if s := x.String(); s != "{1 9 42 144}" {
		t.Errorf("x.String() got: %s, expected: {1 9 42 144}\n", s)
	}

	if b := x.Has(9); b != true {
		t.Errorf("x.Has(9) got: %v, expected: true\n", b)
	}

	if b := x.Has(123); b != false {
		t.Errorf("x.Has(123) got: %v, expected: false\n", b)
	}
}

func TestIntSet_two(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	if s := fmt.Sprint(&x); s != "{1 9 42 144}" {
		t.Errorf("x.Sprintln(&x) got: %s, expected: {1 9 42 144}\n", s)
	}

	if s := x.String(); s != "{1 9 42 144}" {
		t.Errorf("x.String() got: %s, expected: {1 9 42 144}\n", s)
	}

	if s := fmt.Sprint(x); s != "{[4398046511618 0 65536]}" {
		t.Errorf("x.Sprintln(x) got: %s, expected: {[4398046511618 0 65536]}\n", s)
	}
}
