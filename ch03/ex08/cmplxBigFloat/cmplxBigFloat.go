package cmplxBigFloat

import (
	"math/big"
)

const prec = 256

type Cmplx struct {
	real *big.Float
	imag *big.Float
}

func New(z complex128) *Cmplx {
	return &Cmplx{
		big.NewFloat(real(z)),
		big.NewFloat(imag(z)),
	}
}

func Add(x, y *Cmplx) *Cmplx {
	return &Cmplx{
		new(big.Float).SetPrec(prec).Add(x.real, y.real),
		new(big.Float).SetPrec(prec).Add(x.imag, y.imag),
	}
}

func Sub(x, y *Cmplx) *Cmplx {
	return &Cmplx{
		new(big.Float).SetPrec(prec).Sub(x.real, y.real),
		new(big.Float).SetPrec(prec).Sub(x.imag, y.imag),
	}
}

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

func Abs(x *Cmplx) *big.Float {
	return new(big.Float).SetPrec(prec).Sqrt(
		new(big.Float).SetPrec(prec).Add(
			new(big.Float).SetPrec(prec).Mul(x.real, x.real),
			new(big.Float).SetPrec(prec).Mul(x.imag, x.imag),
		),
	)
}
