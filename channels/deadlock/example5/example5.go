package main

import "fmt"

func main() {
	ch := make(chan int)

	go count(ch)

	for {
		msg, ok := <-ch // open is status of the channel
		if !ok {
			break
		}
		fmt.Println(msg)
	}
}

func count(ch chan int) {
	defer close(ch)
	for i := 1; i <= 5; i++ {
		ch <- i
	}
}

// 1
// 2
// 3
// 4
// 5
