package main

import "fmt"

func main() {
	fmt.Println("Hi :3")

	c := make(chan int)

	// send
	go send(c)

	// receive
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func send(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}
