package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 1)
	ch <- "gophers are amazing!"
	fmt.Println(<-ch)

	fmt.Println("done!")
}
