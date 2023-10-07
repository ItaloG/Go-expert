package main

import "fmt"

func recebe(nome string, hello chan<- string) {
	hello <- nome
}

func let(data <-chan string) {
	fmt.Println(<-data)
}

// Thread 1
func main() {
	hello := make(chan string)
	go recebe("Hello", hello)
	let(hello)
}
