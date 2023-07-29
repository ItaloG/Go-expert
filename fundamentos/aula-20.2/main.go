package main

import (
	"fmt"

	"go-expert/matematica"
)

func main() {
	s := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Fiat"}

	fmt.Println(carro.Andar())
	fmt.Println("Resultado: ", s)
	fmt.Println(matematica.A)
}
