package main

import (
	"testing"
)

func createHash(b byte) [32]byte {
	var hash [32]byte
	for i := range &hash {
		hash[i] = b
	}
	return hash
}

func TestShadiff(t *testing.T) {
	var tests = []struct {
		hash1 [32]byte
		hash2 [32]byte
		want  int
	}{
		{createHash(byte(255)), createHash(byte(255)), 0},
		{createHash(byte(255)), createHash(byte(127)), 32},
		{createHash(byte(10)), createHash(byte(3)), 64},
	}
	for _, test := range tests {
		if got := shadiff(&test.hash1, &test.hash2); got != test.want {
			t.Errorf("shadiff(%b, %b) = %v\n", test.hash1, test.hash2, got)
		}
	}
}
