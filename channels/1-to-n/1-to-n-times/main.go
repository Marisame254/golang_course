package main

import "fmt"

func main() {
	fmt.Println("Hi :3")

	n := 10
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 1000; i++ {
			c <- i
		}
		close(c)
	}()

	for i := 0; i < n; i++ {
		go func() {
			for v := range c {
				fmt.Println(v)
			}
			done <- true
		}()
	}

	for i := 0; i < n; i++ {
		<-done
	}

}
