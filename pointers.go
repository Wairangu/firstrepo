package main

import "fmt"

func incre(p *int) int {
	*p++
	return *p
}

func main() {
	var x int
	x = 10
	y := &x
	i := incre(y)
	fmt.Println(i)
}
