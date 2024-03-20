package main

import "fmt"

func main() {
	ch := make(chan int)

	go count(ch)

	for {
		fmt.Println(<-ch)
	}
}

func count(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
}

// 1
// 2
// 3
// 4
// 5
// fatal error: all goroutines are asleep - deadlock!
