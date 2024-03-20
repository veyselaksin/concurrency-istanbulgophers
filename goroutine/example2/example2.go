package main

import (
	"fmt"
	"time"
)

func main() {
	// Example 2
	go func() {
		println("hello, gopher!")
	}()

	fmt.Println("done!")
	time.Sleep(2 * time.Second)
}
