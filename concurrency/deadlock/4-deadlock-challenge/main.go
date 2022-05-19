package main

import "fmt"

func main() {
	fmt.Println("Hi :3")

	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	for {
		fmt.Println(<-c)
	}
}

// This prints the number, but we have recived this error:
// fatal error: allgorutines are asleep - deadlock!
// Where is the deadlock?
// Why are all gorutines asleep
// How can we fix this?
