package main

import (
	"fmt"
	"myBook/distconv"
)

func main() {
	fmt.Println(distconv.FtoM(1).Units())
	fmt.Println(distconv.MtoF(1).Units())
	fmt.Println(distconv.KgToPounds(5).Units())
	fmt.Println(distconv.PoundsToKg(50).Units())

}
