package main

import "fmt"

func main() {
	fmt.Println("Hi :3")
	c := make(chan int)
	c <- 1
	fmt.Println(<-c)
}

// This results in a deadlock
// Can you determine why?
// And what would you do to fix it?
