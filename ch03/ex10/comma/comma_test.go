package comma

import "testing"

func TestComma(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"12", "12"},
		{"123", "123"},
		{"1234", "1,234"},
	}
	for _, test := range tests {
		got := comma(test.input)
		if got != test.want {
			t.Errorf("comma(%v) = %v\n", test.input, got)
		}
	}
}
