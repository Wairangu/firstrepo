package main

import "fmt"

func main() {
	s := []int{10, 4, 6, 1, 4, 6, 7}
	a := [7]int{10, 4, 6, 1, 4, 6, 7}
	fmt.Println("The Array", a)
	fmt.Println("The slice", s)

	fmt.Println("The first element:", s[0])
	fmt.Println("Test", s[:4])

	for i, v := range s {
		fmt.Printf("%d,%d\n", i, v)
	}
}
