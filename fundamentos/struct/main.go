package main

import (
	"fmt"
)

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	italo := Cliente{
		Nome:  "ITalo",
		Idade: 30,
		Ativo: true,
	}
	italo.Ativo = false

	fmt.Println(italo.Ativo)
}
