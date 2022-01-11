package intset

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func randomNumber(rng *rand.Rand, size int) int {
	return rng.Intn(size)
}

func benchmarkIntSetAdd(b *testing.B, size int) {
	seed := time.Now().UTC().UnixNano()
	// b.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	var s IntSet
	for i := 0; i < b.N; i++ {
		s.Add(randomNumber(rng, size))
	}
}

func benchmarkIntSetUnionWith(b *testing.B, size int) {
	seed := time.Now().UTC().UnixNano()
	// b.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	var s1, s2 IntSet
	for i := 0; i < b.N; i++ {
		for j := 0; j < 5; j++ {
			s1.Add(randomNumber(rng, size))
			s2.Add(randomNumber(rng, size))
		}
		s1.UnionWith(&s2)
	}
}

func BenchmarkIntSetAdd100(b *testing.B)        { benchmarkIntSetAdd(b, 100) }
func BenchmarkIntSetAdd10000(b *testing.B)      { benchmarkIntSetAdd(b, 10000) }
func BenchmarkIntSetAdd1000000(b *testing.B)    { benchmarkIntSetAdd(b, 1000000) }
func BenchmarkIntSetAdd10000000(b *testing.B)   { benchmarkIntSetAdd(b, 10000000) }
func BenchmarkIntSetAdd1000000000(b *testing.B) { benchmarkIntSetAdd(b, 1000000000) }

func BenchmarkIntSetUnionWith100(b *testing.B)        { benchmarkIntSetUnionWith(b, 100) }
func BenchmarkIntSetUnionWith10000(b *testing.B)      { benchmarkIntSetUnionWith(b, 10000) }
func BenchmarkIntSetUnionWith1000000(b *testing.B)    { benchmarkIntSetUnionWith(b, 1000000) }
func BenchmarkIntSetUnionWith10000000(b *testing.B)   { benchmarkIntSetUnionWith(b, 10000000) }
func BenchmarkIntSetUnionWith1000000000(b *testing.B) { benchmarkIntSetUnionWith(b, 1000000000) }

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
