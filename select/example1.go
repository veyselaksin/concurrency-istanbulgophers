package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		ch1 <- "hello"

		close(ch1)
	}()

	go func() {
		defer wg.Done()
		ch2 <- "world"

		close(ch2)

	}()

	wg.Wait()
	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	default:
		fmt.Println("default")
	}

}
