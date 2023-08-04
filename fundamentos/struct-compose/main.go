package main

import (
	"fmt"
)

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {
	italo := Cliente{
		Nome:  "ITalo",
		Idade: 30,
		Ativo: true,
	}
	// italo.Endereco.Cidade = "cidade"
	italo.Cidade = "cidade"

	fmt.Println(italo.Ativo)
}
