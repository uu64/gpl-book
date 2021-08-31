package tempconv

import "testing"

func TestCelsiusString(t *testing.T) {
	var tests = []struct {
		input int
		want  string
	}{
		{0, "0°C"},
	}
	for _, test := range tests {
		if got := Celsius(test.input).String(); got != test.want {
			t.Errorf("Celsius(%d).String() = %v", test.input, got)
		}
	}
}

func TestFahrenheitString(t *testing.T) {
	var tests = []struct {
		input int
		want  string
	}{
		{0, "0°F"},
	}
	for _, test := range tests {
		if got := Fahrenheit(test.input).String(); got != test.want {
			t.Errorf("Fahrenheit(%d).String() = %v", test.input, got)
		}
	}
}

func TestKelvinString(t *testing.T) {
	var tests = []struct {
		input int
		want  string
	}{
		{0, "0°K"},
	}
	for _, test := range tests {
		if got := Kelvin(test.input).String(); got != test.want {
			t.Errorf("Kelvin(%d).String() = %v", test.input, got)
		}
	}
}
