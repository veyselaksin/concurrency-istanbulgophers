package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// fanIn function merges data from given channels into one channel.
func fanIn(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(len(channels))

	// Goroutines listening for data from each channel.
	for _, c := range channels {
		go func(c <-chan int) {
			defer wg.Done()
			for n := range c {
				ch <- n
			}
		}(c)
	}

	// Close the ch channel after all goroutines are done.
	go func() {
		defer close(ch)
		wg.Wait()
	}()

	return ch
}

// producer function generates random numbers and writes them to a channel.
func producer() <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < 5; i++ {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			out <- rand.Intn(100)
		}
	}()
	return out
}

func main() {
	// Creating producer channels and merging data using fan-in pattern.
	input1 := producer()
	input2 := producer()
	input3 := producer()

	merged := fanIn(input1, input2, input3)

	// Reading merged data and printing to console.
	for num := range merged {
		fmt.Println(num)
	}
}
