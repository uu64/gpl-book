package cmplxbigfloat

import (
	"math/big"
)

const prec = 256

// Cmplx is complex number using math/big.Float
type Cmplx struct {
	real *big.Float
	imag *big.Float
}

// New creates a new Cmplx object and returns the pointer
func New(z complex128) *Cmplx {
	return &Cmplx{
		big.NewFloat(real(z)),
		big.NewFloat(imag(z)),
	}
}

// Add returns the result of adding two Cmplx numbers
func Add(x, y *Cmplx) *Cmplx {
	return &Cmplx{
		new(big.Float).SetPrec(prec).Add(x.real, y.real),
		new(big.Float).SetPrec(prec).Add(x.imag, y.imag),
	}
}

// Sub returns the result of subtraction of two Cmplx numbers
func Sub(x, y *Cmplx) *Cmplx {
	return &Cmplx{
		new(big.Float).SetPrec(prec).Sub(x.real, y.real),
		new(big.Float).SetPrec(prec).Sub(x.imag, y.imag),
	}
}

// Mul returns the result of multiplying two Cmplx numbers
func Mul(x, y *Cmplx) *Cmplx {
	return &Cmplx{
		new(big.Float).SetPrec(prec).Sub(
			new(big.Float).SetPrec(prec).Mul(x.real, y.real),
			new(big.Float).SetPrec(prec).Mul(x.imag, y.imag),
		),
		new(big.Float).SetPrec(prec).Add(
			new(big.Float).SetPrec(prec).Mul(x.real, y.imag),
			new(big.Float).SetPrec(prec).Mul(x.imag, y.real),
		),
	}
}

// Abs returns the absolute value of a Cmplx number
func Abs(x *Cmplx) *big.Float {
	return new(big.Float).SetPrec(prec).Sqrt(
		new(big.Float).SetPrec(prec).Add(
			new(big.Float).SetPrec(prec).Mul(x.real, x.real),
			new(big.Float).SetPrec(prec).Mul(x.imag, x.imag),
		),
	)
}
