package main

import "fmt"

func main() {
	fmt.Println("Hi :3")

	c := make(chan int)
	go func() {
		c <- 1
	}()

	fmt.Println(<-c)
}
