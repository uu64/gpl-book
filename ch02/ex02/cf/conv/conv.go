package conv

import "fmt"

// Meter unit
type Meter float64

// Feet unit
type Feet float64

// Kilogram unit
type Kilogram float64

// Pound unit
type Pound float64

// String returns a string with Meter unit
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }

// String returns a string with Feet unit
func (f Feet) String() string { return fmt.Sprintf("%gft", f) }

// String returns a string with Kilogram unit
func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }

// String returns a string with Kelvin unit
func (p Pound) String() string { return fmt.Sprintf("%glb", p) }

// MToF converts Meter to Feet
func MToF(m Meter) Feet { return Feet(m * 3.28084) }

// FToM converts Feet to Meter
func FToM(f Feet) Meter { return Meter(f / 3.28084) }

// KToP converts Kilogram to Pound
func KToP(k Kilogram) Pound { return Pound(k * 2.2046) }

// PToK converts Pound to Kilogram
func PToK(p Pound) Kilogram { return Kilogram(p / 2.2046) }
