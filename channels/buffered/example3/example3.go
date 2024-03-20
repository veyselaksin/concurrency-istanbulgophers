package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 1)
	ch <- "gophers are amazing!"
	ch <- "gopherhood!"
	fmt.Println(<-ch)

	fmt.Println("done!")
}

// fatal error: all goroutines are asleep - deadlock!
