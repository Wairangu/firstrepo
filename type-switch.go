package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case int:
		fmt.Printf("Twice %v is %v\n", v, 2*v)
	default:
		fmt.Printf("I do not know the type: %T\n", v)
	}
}

func main() {
	do(21)
	do("Isaac")
	do(3.134543)
	do(true)
}
