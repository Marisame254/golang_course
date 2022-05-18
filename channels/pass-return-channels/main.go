package main

import "fmt"

func main() {
	fmt.Println("Hi :3")

	c := incrementor()
	cSum := puller(c)
	for v := range cSum {
		fmt.Println(v)
	}
}

func incrementor() chan int {
	out := make(chan int)
	go func() {
		count := 10 // change
		for i := 0; i < count; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for v := range c {
			sum += v
		}
		out <- sum
		close(out)
	}()
	return out
}
