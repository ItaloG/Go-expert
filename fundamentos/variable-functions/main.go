package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1, 3, 45, 56, 13))
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
