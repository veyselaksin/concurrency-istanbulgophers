package main

import "fmt"

func main() {
	var ch = make(chan string)

	go func() {
		fmt.Println("First goroutine took the message: ", <-ch)
	}()

	go func() {
		fmt.Println("Second goroutine took the message: ", <-ch)
	}()

	ch <- "gopherhood!"
	fmt.Println("done!")
}
