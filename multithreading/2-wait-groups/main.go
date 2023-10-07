package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}

// Thread 1
func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25)
	// Tread 2
	go task("A", &waitGroup)
	// Tread 3
	go task("B", &waitGroup)
	// Tread 4
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
			waitGroup.Done()
		}
	}()

	waitGroup.Wait()
}
