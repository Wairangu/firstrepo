package main

import "fmt"

func fToC(f float64) float64 {
	c := (f - 32) * 5 / 9
	return c
}

func cToF(c float64) float64 {
	f := 32 + (9 / 5 * c)
	return f
}

func main() {
	const boilingC, boilingF, freezingC, freezingF = 100.0, 212.0, 0.0, 32.0
	b := cToF(boilingC)
	d := fToC(boilingF)
	fc := fToC(freezingF)
	ff := cToF(freezingC)
	fmt.Println("Boiling point in F", b)
	fmt.Println("Boiling point in C", d)
	fmt.Println("Freezing point in C", fc)
	fmt.Println("Freezing point in F", ff)
}
