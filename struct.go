package main

import "fmt"

//Vertex holds the measurements for vertices
type Vertex struct {
	X float64
	Y float64
}

func main() {
	v := Vertex{2.4, 5.6}
	fmt.Println(v)
	fmt.Println("X is", v.X)
	fmt.Println("Y is", v.Y)
	v.Y = 12.6
	fmt.Println("New Y is", v.Y)
	p := &v
	fmt.Println("Pointer ", (*p).Y)
	p.X = 55.6
	fmt.Println("New Xx value is", v.X)
}
