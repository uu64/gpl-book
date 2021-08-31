package tempconv

import "fmt"

type Celsius float64

type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// String returns a string with units
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

// String returns a string with units
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
