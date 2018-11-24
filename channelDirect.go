package main

import "fmt"

// only accepts chan for sending
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// accepts one channel for receive
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "Passed message.")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
