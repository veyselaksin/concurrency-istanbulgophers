package main

import (
	"fmt"
)

func main() {
	var ch = make(chan string)

	go func() {
		ch <- "hello gophers!" // write to channel, blocks until data is read
	}()

	msg := <-ch // read from channel, blocks until data is available
	fmt.Println(msg)
}

// hello gophers!
