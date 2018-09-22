/*This Go package converts temperatures among Celsius, Fahrenheit and Kelvin
scales. */

package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

//functions to indicate the scale for each temp scale
func (c Celsius) Scale() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) Scale() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) Scale() string     { return fmt.Sprintf("%gK", k) }

//function to convert from F to C
func FtoC(f Fahrenheit) Celsius {
	r := (f - 32) * 5 / 9
	return Celsius(r)
}

//function to convert from C to F
func CtoF(c Celsius) Fahrenheit {
	y := 32 + c * 1.8
	return Fahrenheit(y)
}

//faction to convert from C to F
func CtoK(c Celsius) Kelvin {
	return Kelvin(c) + Kelvin(273.15)
}
