package main

import "fmt"

/**
A palavra defer faz com que uma linha
de código seja excutada no fim do processo
*/

func main() {
	defer fmt.Println("Primeira Linha")
	fmt.Println("Segunda Linha")
	fmt.Println("Terceira Linha")
}
