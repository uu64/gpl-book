package intset

import (
	"testing"
)

func TestLen(t *testing.T) {
	var x IntSet
	if got := x.Len(); got != 0 {
		t.Errorf("(%v).Len() returns %v\n", x.String(), got)
	}

	x.Add(1)
	x.Add(9)
	if got := x.Len(); got != 2 {
		t.Errorf("(%v).Len() returns %v\n", x.String(), got)
	}

	for i := 0; i < 64; i++ {
		x.Add(i)
	}
	if got := x.Len(); got != 64 {
		t.Errorf("(%v).Len() returns %v\n", x.String(), got)
	}

	for i := 0; i < 100; i++ {
		x.Add(i)
	}
	if got := x.Len(); got != 100 {
		t.Errorf("(%v).Len() returns %v\n", x.String(), got)
	}
}

func TestRemove(t *testing.T) {
	var x IntSet
	x.Remove(9)
	if got := x.Len(); got != 0 {
		t.Errorf("(%v).Len() returns %v\n", x.String(), got)
	}

	x.Add(1)
	x.Add(144)
	x.Add(9)

	x.Remove(9)
	if got := x.Len(); got != 2 {
		t.Errorf("(%v).Len() returns %v\n", x.String(), got)
	}
	if !(x.Has(1) && x.Has(144) && !x.Has(9)) {
		t.Errorf("remove error: want {1 144}, got %v\n", x.String())
	}

	x.Remove(2)
	if got := x.Len(); got != 2 {
		t.Errorf("(%v).Len() returns %v\n", x.String(), got)
	}
	if !(x.Has(1) && x.Has(144) && !x.Has(9)) {
		t.Errorf("remove error: want {1 144}, got %v\n", x.String())
	}
}

func TestClear(t *testing.T) {
	var x IntSet
	x.Clear()
	if got := x.Len(); got != 0 {
		t.Errorf("(%v).Len() returns %v\n", x.String(), got)
	}

	x.Add(1)
	x.Add(9)
	x.Clear()
	if got := x.Len(); got != 0 {
		t.Errorf("(%v).Len() returns %v\n", x.String(), got)
	}

	for i := 0; i < 64; i++ {
		x.Add(i)
	}
	x.Clear()
	if got := x.Len(); got != 0 {
		t.Errorf("(%v).Len() returns %v\n", x.String(), got)
	}

	for i := 0; i < 100; i++ {
		x.Add(i)
	}
	x.Clear()
	if got := x.Len(); got != 0 {
		t.Errorf("(%v).Len() returns %v\n", x.String(), got)
	}
}

func TestCopy(t *testing.T) {
	var x IntSet

	x.Add(1)
	x.Add(144)
	x.Add(9)
	y := x.Copy()
	if got := y.Len(); got != 3 {
		t.Errorf("copy error: want 3, got %v\n", got)
	}
	if !y.Has(1) || !y.Has(9) || !y.Has(144) {
		t.Errorf("copy error: want {1 9 144}, got %v\n", y.String())
	}

	x.Remove(9)
	if !y.Has(1) || !y.Has(9) || !y.Has(144) {
		t.Errorf("copy error: want {1 9 144}, got %v\n", y.String())
	}
}
