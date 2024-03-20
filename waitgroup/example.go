package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done() // Decrease the counter of the WaitGroup
		fmt.Println("First goroutine")
	}()

	go func() {
		defer wg.Done() // Decrease the counter of the WaitGroup
		fmt.Println("Second goroutine")
	}()

	wg.Wait() // Wait until the counter of the WaitGroup is zero.
	fmt.Println("done!")
}
