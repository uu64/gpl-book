package main

import "testing"

func TestJoin(t *testing.T) {
	tests := []struct {
		sep   string
		elems []string
		want  string
	}{
		{",", []string{}, ""},
		{",", []string{"a"}, "a"},
		{" ", []string{"a", "b"}, "a b"},
		{"", []string{"a", "b"}, "ab"},
		{"///", []string{"a", "b", "c"}, "a///b///c"},
	}
	for _, test := range tests {
		if got := join(test.sep, test.elems...); got != test.want {
			t.Errorf("error: join(%s, %s) returns %v\n", test.sep, test.elems, got)
		}
	}
}
