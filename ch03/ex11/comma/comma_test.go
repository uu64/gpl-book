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
		{"12.", "12."},
		{"123.", "123."},
		{"1234.", "1,234."},
		{"56.1", "56.1"},
		{"567.1", "567.1"},
		{"5678.1", "5,678.1"},
		{"+56.12", "+56.12"},
		{"+567.12", "+567.12"},
		{"+5678.12", "+5,678.12"},
		{"-12.000", "-12.000"},
		{"-123.000", "-123.000"},
		{"-1234.000", "-1,234.000"},
	}
	for _, test := range tests {
		got := comma(test.input)
		if got != test.want {
			t.Errorf("comma(%v) = %v\n", test.input, got)
		}
	}
}
