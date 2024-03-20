package main

import "time"

func send(ch chan string, msg string) {
	ch <- msg
}

func recv(ch chan string) string {
	return <-ch
}

func main() {
	var ch = make(chan string)

	messages := []string{"hello", "world", "from", "channels"}

	for _, msg := range messages {
		go send(ch, msg) // send is a blocking operation
	}

	for range messages {
		println(recv(ch))
		time.Sleep(1 * time.Second)
	}

}

// channels
// world
// hello
// from
