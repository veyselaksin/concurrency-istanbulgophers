package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	done := make(chan struct{}) // signal channel

	go func() {
		fmt.Println(<-ch)
		done <- struct{}{} // send signal
	}()

	ch <- "I love golang!" // write to channel, blocks until data is read
	<-done

	fmt.Println("done!")
}
