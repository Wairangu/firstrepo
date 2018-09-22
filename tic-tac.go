package main

import (
	"fmt"
	"strings"
)

func main() {
	s := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	s[0][0] = "0"

	for i := 0; i < len(s); i++ {
		s[i][i] = "X"
	}

	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", strings.Join(s[i], " "))
	}
}
