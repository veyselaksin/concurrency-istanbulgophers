package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	ch <- "deadlock!"

	fmt.Println(<-ch)
}
