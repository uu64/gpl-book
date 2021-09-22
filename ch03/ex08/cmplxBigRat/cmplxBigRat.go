package cmplxBigRat

import (
	"math/big"
)

type Cmplx struct {
	real *big.Rat
	imag *big.Rat
}

func New(z complex128) *Cmplx {
	return &Cmplx{
		new(big.Rat).SetFloat64(real(z)),
		new(big.Rat).SetFloat64(imag(z)),
	}
}

func Add(x, y *Cmplx) *Cmplx {
	return &Cmplx{
		new(big.Rat).Add(x.real, y.real),
		new(big.Rat).Add(x.imag, y.imag),
	}
}

func Sub(x, y *Cmplx) *Cmplx {
	return &Cmplx{
		new(big.Rat).Sub(x.real, y.real),
		new(big.Rat).Sub(x.imag, y.imag),
	}
}

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

func SqAbs(x *Cmplx) *big.Rat {
	return new(big.Rat).Add(
		new(big.Rat).Mul(x.real, x.real),
		new(big.Rat).Mul(x.imag, x.imag),
	)
}
