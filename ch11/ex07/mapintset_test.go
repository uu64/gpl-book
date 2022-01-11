package intset

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func benchmarkMapIntSetAdd(b *testing.B, size int) {
	seed := time.Now().UTC().UnixNano()
	// b.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	var s IntSet
	for i := 0; i < b.N; i++ {
		s.Add(randomNumber(rng, size))
	}
}

func benchmarkMapIntSetUnionWith(b *testing.B, size int) {
	seed := time.Now().UTC().UnixNano()
	// b.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	s1, s2 := make(MapIntSet), make(MapIntSet)
	for i := 0; i < b.N; i++ {
		for j := 0; j < 5; j++ {
			s1.Add(randomNumber(rng, size))
			s2.Add(randomNumber(rng, size))
		}
		s1.UnionWith(s2)
	}
}

func BenchmarkMapIntSetAdd100(b *testing.B)         { benchmarkMapIntSetAdd(b, 100) }
func BenchmarkMapIntSetAdd10000(b *testing.B)       { benchmarkMapIntSetAdd(b, 10000) }
func BenchmarkMapIntSetAdd1000000(b *testing.B)     { benchmarkMapIntSetAdd(b, 1000000) }
func BenchmarkMapIntSetAdd100000000(b *testing.B)   { benchmarkMapIntSetAdd(b, 100000000) }
func BenchmarkMapIntSetAdd10000000000(b *testing.B) { benchmarkMapIntSetAdd(b, 10000000000) }

func BenchmarkMapIntSetUnionWith100(b *testing.B)        { benchmarkMapIntSetUnionWith(b, 100) }
func BenchmarkMapIntSetUnionWith10000(b *testing.B)      { benchmarkMapIntSetUnionWith(b, 10000) }
func BenchmarkMapIntSetUnionWith1000000(b *testing.B)    { benchmarkMapIntSetUnionWith(b, 1000000) }
func BenchmarkMapIntSetUnionWith10000000(b *testing.B)   { benchmarkMapIntSetUnionWith(b, 10000000) }
func BenchmarkMapIntSetUnionWith1000000000(b *testing.B) { benchmarkMapIntSetUnionWith(b, 1000000000) }

func TestMapIntSet_one(t *testing.T) {
	x, y := make(MapIntSet), make(MapIntSet)
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

	x.UnionWith(y)
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

func TestMapIntSet_two(t *testing.T) {
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
