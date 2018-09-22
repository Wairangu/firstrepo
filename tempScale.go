package main

import (
	"fmt"
	"myBook/tempconv"
)

func main() {
	fmt.Println(tempconv.BoilingC)

	fmt.Println(tempconv.FtoC(100).Scale())
	fmt.Println(tempconv.CtoF(100).Scale())
	fmt.Println(tempconv.CtoK(100).Scale())

}
