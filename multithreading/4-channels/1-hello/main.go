package main

import "fmt"

// Thread 1
func main() {
	canal := make(chan string) // Vazio

	// Thread 2
	go func() {
		canal <- "Óla mundo" // Está cheio
	}()

	// Thread 1
	msg := <-canal // Canal ezvazia
	fmt.Println(msg)
}
