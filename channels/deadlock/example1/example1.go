package main

func sleepGopher(ch chan string) string {
	return <-ch
}

func main() {
	var ch = make(chan string)

	msg := sleepGopher(ch)
	println(msg)
}
