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
		t.Errorf("Copy error: length should be 3, but %v\n", got)
	}
	if !y.Has(1) || !y.Has(9) || !y.Has(144) {
		t.Errorf("Copy error: want {1 9 144}, got %v\n", y.String())
	}

	x.Remove(9)
	if !y.Has(1) || !y.Has(9) || !y.Has(144) {
		t.Errorf("Copy error: want {1 9 144}, got %v\n", y.String())
	}
}

func TestAddAll(t *testing.T) {
	var x IntSet

	x.AddAll(1, 144, 9)
	if got := x.Len(); got != 3 {
		t.Errorf("AddAll error: length should be 3, but %v\n", got)
	}
	if !x.Has(1) || !x.Has(9) || !x.Has(144) {
		t.Errorf("AddAll error: want {1 9 144}, got %v\n", x.String())
	}
}

func TestIntersectWith(t *testing.T) {
	var x, y IntSet

	y.AddAll(1, 144, 9)
	x.IntersectWith(&y)
	if got := x.Len(); got != 0 {
		t.Errorf("IntersectWith error: length should be 0, but %v\n", got)
	}
	x.Clear()
	y.Clear()

	x.AddAll(1, 144, 9)
	y.AddAll(144)
	x.IntersectWith(&y)
	if got := x.Len(); got != 1 {
		t.Errorf("IntersectWith error: length should be 1, but %v\n", got)
	}
	if !x.Has(144) {
		t.Errorf("IntersectWith error: want {144}, got %v\n", x.String())
	}
	x.Clear()
	y.Clear()

	x.AddAll(1, 144, 9)
	y.AddAll(144, 1, 56)
	x.IntersectWith(&y)
	if got := x.Len(); got != 2 {
		t.Errorf("IntersectWith error: length should be 2, but %v\n", got)
	}
	if !(x.Has(1) && x.Has(144)) {
		t.Errorf("IntersectWith error: want {1 144}, got %v\n", x.String())
	}
	x.Clear()
	y.Clear()
}

func TestDifferenceWith(t *testing.T) {
	var x, y IntSet

	y.AddAll(1, 144, 9)
	x.DifferenceWith(&y)
	if got := x.Len(); got != 0 {
		t.Errorf("DifferenceWith error: length should be 0, but %v\n", got)
	}
	x.Clear()
	y.Clear()

	x.AddAll(1, 144, 9)
	y.AddAll(144)
	x.DifferenceWith(&y)
	if got := x.Len(); got != 2 {
		t.Errorf("DifferenceWith error: length should be 2, but %v\n", got)
	}
	if !(x.Has(1) && x.Has(9)) {
		t.Errorf("DifferenceWith error: want {1 9}, got %v\n", x.String())
	}
	x.Clear()
	y.Clear()

	x.AddAll(1, 144, 9)
	y.AddAll(144, 1, 56)
	x.DifferenceWith(&y)
	if got := x.Len(); got != 1 {
		t.Errorf("DifferenceWith error: length should be 1, but %v\n", got)
	}
	if !x.Has(9) {
		t.Errorf("DifferenceWith error: want {9}, got %v\n", x.String())
	}
	x.Clear()
	y.Clear()

	x.AddAll(1, 144, 9)
	y.AddAll(144, 1, 9)
	x.DifferenceWith(&y)
	if got := x.Len(); got != 0 {
		t.Errorf("DifferenceWith error: length should be 0, but %v\n", got)
	}
	x.Clear()
	y.Clear()
}
func TestSymmetricDifferenceWith(t *testing.T) {
	var x, y IntSet

	y.AddAll(1, 144, 9)
	x.SymmetricDifferenceWith(&y)
	if got := x.Len(); got != 3 {
		t.Errorf("SymmetricDifferenceWith error: length should be 3, but %v\n", got)
	}
	x.Clear()
	y.Clear()

	x.AddAll(1, 144, 9)
	y.AddAll(144)
	x.SymmetricDifferenceWith(&y)
	if got := x.Len(); got != 2 {
		t.Errorf("SymmetricDifferenceWith error: length should be 2, but %v\n", got)
	}
	if !(x.Has(1) && x.Has(9)) {
		t.Errorf("SymmetricDifferenceWith error: want {1 9}, got %v\n", x.String())
	}
	x.Clear()
	y.Clear()

	x.AddAll(1, 144, 9)
	y.AddAll(144, 1, 56)
	x.SymmetricDifferenceWith(&y)
	if got := x.Len(); got != 2 {
		t.Errorf("SymmetricDifferenceWith error: length should be 1, but %v\n", got)
	}
	if !(x.Has(9) && x.Has(56)) {
		t.Errorf("SymmetricDifferenceWith error: want {9}, got %v\n", x.String())
	}
	x.Clear()
	y.Clear()

	x.AddAll(1, 144, 9)
	y.AddAll(144, 1, 9)
	x.SymmetricDifferenceWith(&y)
	if got := x.Len(); got != 0 {
		t.Errorf("SymmetricDifferenceWith error: length should be 0, but %v\n", got)
	}
	x.Clear()
	y.Clear()
}
