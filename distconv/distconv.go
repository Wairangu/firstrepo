package distconv

import "fmt"

type Feet float64
type Yards float64
type Metres float64
type Kilo float64
type Pounds float64

func (f Feet) Units() string   { return fmt.Sprintf("%gft", f) }
func (y Yards) Units() string  { return fmt.Sprintf("%gyd", y) }
func (m Metres) Units() string { return fmt.Sprintf("%gM", m) }
func (p Pounds) Units() string { return fmt.Sprintf("%glb", p) }
func (k Kilo) Units() string { return fmt.Sprintf("%gKG", k) }
//convert feet to metres
func FtoM( f Feet) Metres{
	return Metres(f * 0.3048)
}
//convert metres to feet
func MtoF(m Metres) Feet  {
	return Feet(m * 3.28084)
}
//Kg to Pounds
func KgToPounds( k Kilo) Pounds  {
	return Pounds(k * 2.2046)
}

//Pponds to KG
func PoundsToKg( p Pounds) Kilo{
	return Kilo(p * 0.45359237)
}


