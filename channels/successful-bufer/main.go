package main

import "fmt"

func main() {
	fmt.Println("Hi :3")

	// channels block =========
	// c := make(chan int)
	// c <- 42
	// fmt.Println(<-c)
	// =========================

	c := make(chan int, 2)
	go func() {
		c <- 42
		c <- 43
	}()
	fmt.Println(<-c)
	fmt.Println(<-c)

}
