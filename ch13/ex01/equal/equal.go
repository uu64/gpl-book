package equal

import (
	"fmt"
	"math"
	"reflect"
)

const threshold = 0.000000001

func Equal(x, y interface{}) bool {
	return equal(reflect.ValueOf(x), reflect.ValueOf(y))
}

func equal(x, y reflect.Value) bool {
	if x.Type() != y.Type() {
		return false
	}

	switch x.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64:
		return x.Int() == y.Int()

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr:
		return x.Uint() == y.Uint()

	case reflect.Float32, reflect.Float64:
		d := math.Abs(x.Float() - y.Float())
		return d < threshold

	case reflect.Complex64, reflect.Complex128:
		xRe := real(x.Complex())
		xIm := imag(x.Complex())
		yRe := real(y.Complex())
		yIm := imag(y.Complex())
		re := xRe - yRe
		im := xIm - yIm
		return math.Sqrt(re*re+im*im) < threshold

	case reflect.Ptr, reflect.Interface:
		return equal(x.Elem(), y.Elem())

	default:
		panic(fmt.Sprintf("unsupported type: %s", x.Kind()))
	}
}
