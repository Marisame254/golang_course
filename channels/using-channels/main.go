package main

import "fmt"

func main() {
	fmt.Println("Hi :3")

	c := make(chan int)

	// send
	go send(c)

	// receive
	recieve(c)

	fmt.Println("about to exit")
}

func send(c chan<- int) {
	c <- 31
}

func recieve(c <-chan int) {
	fmt.Println(<-c)
}
