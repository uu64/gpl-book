package tempconv

import (
	"flag"
	"fmt"
)

// Celsius unit
type Celsius float64

// Fahrenheit unit
type Fahrenheit float64

// Kelvin unit
type Kelvin float64

// CToF converts Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32.0) }

// FToC converts Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }

// KToC converts Kelvin to Celsius
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// String implements the flag.Value
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
