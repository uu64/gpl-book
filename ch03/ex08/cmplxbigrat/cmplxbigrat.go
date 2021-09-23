package cmplxbigrat

import (
	"math/big"
)

// Cmplx is complex number using math/big.Rat
type Cmplx struct {
	real *big.Rat
	imag *big.Rat
}

// New creates a new Cmplx object and returns the pointer
func New(z complex128) *Cmplx {
	return &Cmplx{
		new(big.Rat).SetFloat64(real(z)),
		new(big.Rat).SetFloat64(imag(z)),
	}
}

// Add returns the result of adding two Cmplx numbers
func Add(x, y *Cmplx) *Cmplx {
	return &Cmplx{
		new(big.Rat).Add(x.real, y.real),
		new(big.Rat).Add(x.imag, y.imag),
	}
}

// Sub returns the result of subtraction of two Cmplx numbers
func Sub(x, y *Cmplx) *Cmplx {
	return &Cmplx{
		new(big.Rat).Sub(x.real, y.real),
		new(big.Rat).Sub(x.imag, y.imag),
	}
}

// Mul returns the result of multiplying two Cmplx numbers
func Mul(x, y *Cmplx) *Cmplx {
	return &Cmplx{
		new(big.Rat).Sub(
			new(big.Rat).Mul(x.real, y.real),
			new(big.Rat).Mul(x.imag, y.imag),
		),
		new(big.Rat).Add(
			new(big.Rat).Mul(x.real, y.imag),
			new(big.Rat).Mul(x.imag, y.real),
		),
	}
}

// SqAbs returns the square of the absolute value of a Cmplx number
func SqAbs(x *Cmplx) *big.Rat {
	return new(big.Rat).Add(
		new(big.Rat).Mul(x.real, x.real),
		new(big.Rat).Mul(x.imag, x.imag),
	)
}
