package main

import "fmt"

//Person holds name ans age of a person
type Person struct {
	Name string
	Age  int
}

//Units holds units for various units
type Units struct {
	Value float64
	Unit  string
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years old)", p.Name, p.Age)
}

func (i Units) String() string {
	return fmt.Sprintf("%v %v", i.Value, i.Unit)
}

func main() {
	a := Person{"Isaac Muthui", 30}
	z := Person{"Fatuma Zarika", 26}
	fmt.Println(a, z)

	b := Units{35.4, "meters"}
	c := Units{40, "Inches"}
	fmt.Println(b, c)
}
