// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

//Celsius Temperature
type Celsius float64

// Fahrenheit Temperature
type Fahrenheit float64

// Kelvin Temperature
type Kelvin float64

const (
	// AbsoluteZeroC is the lowest temperature in Celsius
	AbsoluteZeroC Celsius = -273.15
	// FreezingC is the freezing temperature of pure water in Celsius
	FreezingC Celsius = 0
	// BoilingC is the boiling temperature of pure water in Celsius
	BoilingC Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%.2f°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.2f°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%.2f°K", k) }
