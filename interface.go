package main

import (
	"fmt"
	"math"
)

type myFloat float64

//Abser is an interface that has the hypo method
type Abser interface {
	hypo() float64
}

//Sides bolds the dimensions of a triangle
type Sides struct {
	height, breadth float64
}

type dimensions struct {
	side1, side2, side3 float64
}

func (v *Sides) hypo() float64 {
	return math.Sqrt(v.breadth*v.breadth + v.height*v.height)
}

//pointer receiver Scales the triangle by 25
func (v *Sides) scaler(f float64) {
	v.height = v.height * f
	v.breadth = v.breadth * f
}

func (f myFloat) hypo() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (z *dimensions) hypo() float64 {
	return z.side1 + z.side2 + z.side3

}

func main() {
	var a Abser

	f := myFloat(math.Sqrt2)
	v := Sides{3, 4}
	z := dimensions{10, 10, 10}
	v.scaler(10) //scales the sides by 10
	a = f
	fmt.Println("Using a=f", a.hypo())
	a = &v
	fmt.Println("Using a=value of pointer to Sides", a.hypo())
	a = &z
	fmt.Println("Perimeter of a triangle", a.hypo())

	//

}
