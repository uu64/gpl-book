package conv

import (
	"testing"
)

func TestMToF(t *testing.T) {
	var tests = []struct {
		input Meter
		want  Feet
	}{
		{Meter(1), Feet(3.28084)},
	}
	for _, test := range tests {
		if got := MToF(test.input); got != test.want {
			t.Errorf("MtoF(%v) = %v", test.input, got)
		}
	}
}

func TestFToM(t *testing.T) {
	var tests = []struct {
		input Feet
		want  Meter
	}{
		{Feet(3.28084), Meter(1)},
	}
	for _, test := range tests {
		if got := FToM(test.input); got != test.want {
			t.Errorf("FtoM(%v) = %v", test.input, got)
		}
	}
}

func TestKtoP(t *testing.T) {
	var tests = []struct {
		input Kilogram
		want  Pound
	}{
		{Kilogram(1), Pound(2.2046)},
	}
	for _, test := range tests {
		if got := KToP(test.input); got != test.want {
			t.Errorf("KtoP(%v) = %v", test.input, got)
		}
	}
}

func TestPtoK(t *testing.T) {
	var tests = []struct {
		input Pound
		want  Kilogram
	}{
		{Pound(2.2046), Kilogram(1)},
	}
	for _, test := range tests {
		if got := PToK(test.input); got != test.want {
			t.Errorf("PtoK(%v) = %v", test.input, got)
		}
	}
}
