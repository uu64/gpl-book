package tempconv

import "fmt"

// Celsius unit
type Celsius float64

// Fahrenheit unit
type Fahrenheit float64

// Kelvin unit
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	AbsoluteZeroK Kelvin  = 0
	FreezingK     Kelvin  = 273.15
	BoilingK      Kelvin  = 373.15
)

// String returns a string with Celsius unit
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

// String returns a string with Fahrenheit unit
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

// String returns a string with Kelvin unit
func (k Kelvin) String() string { return fmt.Sprintf("%g°K", k) }
