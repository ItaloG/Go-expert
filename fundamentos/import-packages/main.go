package main

import "fmt"

type ID int

var (
	id ID = 1
)

func main() {
	fmt.Printf("O tipo de id é %T", id)
}
