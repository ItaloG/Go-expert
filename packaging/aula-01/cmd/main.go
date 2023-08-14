package main

import (
	"fmt"

	"github.com/ItaloG/Go-expert/packaging/aula-01/math"
)

func main() {
	m := math.NewMath(1, 2)
	m.C = 3
	fmt.Println(m.C)
	// fmt.Println(m.Add())
	// fmt.Println(math.X)
}
