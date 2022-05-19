package main

import "fmt"

func main() {
	fmt.Println("Hi :3")

	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}

// remenber to close your channel
// if you do not close tour channel, you will receive this error
// fatal error: all gorutines are asleddp - deadlock!

// ************** IMPORTANT ***************
// YOU NEED GO VERSION 1.5.2 OR GREATER
// otherwise you will receive this error
// fatal error: all gorutines are asleep - deadlock!
