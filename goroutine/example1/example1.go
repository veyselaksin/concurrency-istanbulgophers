package main

import (
	"fmt"
)

func main() {

	go func() {
		println("hello, gopher!")
	}()

	fmt.Println("done!")
}
