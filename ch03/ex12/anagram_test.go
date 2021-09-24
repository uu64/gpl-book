package main

import "testing"

func TestIsAnagram(t *testing.T) {
	var tests = []struct {
		s1   string
		s2   string
		want bool
	}{
		{"canoe", "ocean", true},
		{"Cinerama", "American", true},
		{"anagrams", "ARS MAGNA", true},
		{"Statue of Liberty", "built to stay free", true},
		{"one plus twelve", "two plus eleven", true},
		{"あいうえお", "うえおあい", true},
		{"apple", "banana", false},
		{"hello!", "Hello!!", false},
		{"あいうえお", "あかさたな", false},
	}
	for _, test := range tests {
		got := isAnagram(test.s1, test.s2)
		if got != test.want {
			t.Errorf("isAnagram(%v, %v) = %v\n", test.s1, test.s2, got)
		}
	}
}
