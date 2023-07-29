package main

import "fmt"

func main() {
	salarios := map[string]int{"ITalo": 1000, "Jão": 2000, "Maria": 3000}
	// sal := make(map[string]int)
	// sal1 := map[string]int{}

	delete(salarios, "Italo")
	salarios["Gabriel"] = 5000

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}

}
