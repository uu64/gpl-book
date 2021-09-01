package tempconv

import (
	"testing"
)

func TestCToF(t *testing.T) {
	var tests = []struct {
		input Celsius
		want  Fahrenheit
	}{
		{BoilingC, Fahrenheit(212)},
	}
	for _, test := range tests {
		if got := CToF(test.input); got != test.want {
			t.Errorf("CtoF(%v) = %v", test.input, got)
		}
	}
}

func TestCToK(t *testing.T) {
	var tests = []struct {
		input Celsius
		want  Kelvin
	}{
		{AbsoluteZeroC, Kelvin(0)},
	}
	for _, test := range tests {
		if got := CToK(test.input); got != test.want {
			t.Errorf("CtoK(%v) = %v", test.input, got)
		}
	}
}

func TestFToC(t *testing.T) {
	var tests = []struct {
		input Fahrenheit
		want  Celsius
	}{
		{Fahrenheit(212), BoilingC},
	}
	for _, test := range tests {
		if got := FToC(test.input); got != test.want {
			t.Errorf("FtoC(%v) = %v", test.input, got)
		}
	}
}

func TestFToK(t *testing.T) {
	var tests = []struct {
		input Fahrenheit
		want  Kelvin
	}{
		{Fahrenheit(-459.67), AbsoluteZeroK},
	}
	for _, test := range tests {
		if got := FToK(test.input); got != test.want {
			t.Errorf("FtoK(%v) = %v", test.input, got)
		}
	}
}

func TestKToC(t *testing.T) {
	var tests = []struct {
		input Kelvin
		want  Celsius
	}{
		{FreezingK, FreezingC},
	}
	for _, test := range tests {
		if got := KToC(test.input); got != test.want {
			t.Errorf("KtoC(%v) = %v", test.input, got)
		}
	}
}

func TestKToF(t *testing.T) {
	var tests = []struct {
		input Kelvin
		want  Fahrenheit
	}{
		{BoilingK, Fahrenheit(212)},
	}
	for _, test := range tests {
		if got := KToF(test.input); got != test.want {
			t.Errorf("KtoF(%v) = %v", test.input, got)
		}
	}
}
