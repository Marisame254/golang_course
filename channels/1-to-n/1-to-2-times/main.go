package main

import "fmt"

func main() {
	fmt.Println("Hi :3")

	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 1000; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		for n := range c {
			fmt.Println(n)
		}
		done <- true
	}()

	go func() {
		for v := range c {
			fmt.Println(v)
		}
		done <- true
	}()

	<-done
	<-done
}
