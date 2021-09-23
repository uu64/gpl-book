package cmplxbigrat

import (
	"fmt"
	"math/big"
	"testing"
)

func TestNew(t *testing.T) {
	var tests = []struct {
		input    complex128
		expected *Cmplx
	}{
		{complex(1, 3), &Cmplx{new(big.Rat).SetFloat64(1), new(big.Rat).SetFloat64(3)}},
	}
	for _, test := range tests {
		got := New(test.input)
		if got.real.Cmp(test.expected.real) != 0 || got.imag.Cmp(test.expected.imag) != 0 {
			t.Errorf("Newcmplxbigrat(%v) = %v\n", test.input, toString(got))
		}
	}
}

func TestAdd(t *testing.T) {
	var tests = []struct {
		input    [2]*Cmplx
		expected *Cmplx
	}{
		{[2]*Cmplx{New(complex(2, 1)), New(complex(-1, 3))}, New(complex(1, 4))},
	}
	for _, test := range tests {
		got := Add(test.input[0], test.input[1])
		if got.real.Cmp(test.expected.real) != 0 || got.imag.Cmp(test.expected.imag) != 0 {
			t.Errorf("Add(%v, %v) = %v\n", toString(test.input[0]), toString(test.input[1]), toString(got))
		}
	}
}

func TestSub(t *testing.T) {
	var tests = []struct {
		input    [2]*Cmplx
		expected *Cmplx
	}{
		{[2]*Cmplx{New(complex(2, 1)), New(complex(-1, 3))}, New(complex(3, -2))},
	}
	for _, test := range tests {
		got := Sub(test.input[0], test.input[1])
		if got.real.Cmp(test.expected.real) != 0 || got.imag.Cmp(test.expected.imag) != 0 {
			t.Errorf("Sub(%v, %v) = %v\n", toString(test.input[0]), toString(test.input[1]), toString(got))
		}
	}
}

func TestMul(t *testing.T) {
	var tests = []struct {
		input    [2]*Cmplx
		expected *Cmplx
	}{
		{[2]*Cmplx{New(complex(2, 1)), New(complex(-1, 3))}, New(complex(-5, 5))},
	}
	for _, test := range tests {
		got := Mul(test.input[0], test.input[1])
		if got.real.Cmp(test.expected.real) != 0 || got.imag.Cmp(test.expected.imag) != 0 {
			t.Errorf("Mul(%v, %v) = %v\n", toString(test.input[0]), toString(test.input[1]), toString(got))
		}
	}
}

func TestAbs(t *testing.T) {
	var tests = []struct {
		input    *Cmplx
		expected *big.Rat
	}{
		{New(complex(3, 4)), new(big.Rat).SetFloat64(25)},
	}
	for _, test := range tests {
		got := SqAbs(test.input)
		if got.Cmp(test.expected) != 0 {
			t.Errorf("Abs(%v) = %v\n", toString(test.input), got)
		}
	}
}

func toString(o *Cmplx) string {
	return fmt.Sprintf("{%v, %v}", o.real, o.imag)
}
