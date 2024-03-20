package main

import (
	"fmt"
	"time"
)

func main() {
	var maxGoroutines = 10
	ch := make(chan int, maxGoroutines) // this is a buffered channel with a capacity of 20

	const maxSemaphore = 5
	sch := make(chan struct{}, maxSemaphore) // this is a buffered channel with a capacity of 5

	for g := 0; g < maxGoroutines; g++ {
		go func(n int) {
			sch <- struct{}{}           // there is no blocking here because the channel has a capacity of 5
			time.Sleep(2 * time.Second) // simulate work

			ch <- n // there is no blocking here because the channel has a capacity of 20
			<-sch   // block until there is capacity in the semaphore
		}(g)
	}

	// read from the channel
	for maxGoroutines > 0 {
		fmt.Println("Max Goroutine: ", maxGoroutines, "Info in Channel: ", <-ch)
		maxGoroutines--
	}

	fmt.Println("done!")
}

// Max Goroutine:  10 Info in Channel:  8
// Max Goroutine:  9 Info in Channel:  7
// Max Goroutine:  8 Info in Channel:  6
// Max Goroutine:  7 Info in Channel:  9
// Max Goroutine:  6 Info in Channel:  4
// Max Goroutine:  5 Info in Channel:  0
// Max Goroutine:  4 Info in Channel:  2
// Max Goroutine:  3 Info in Channel:  1
// Max Goroutine:  2 Info in Channel:  3
// Max Goroutine:  1 Info in Channel:  5
// done!
