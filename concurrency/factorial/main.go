package main

import "fmt"

func main() {
	fmt.Println("Hi :3")

	c := factorial(4)

	for v := range c {
		fmt.Println(v)
	}
}

func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		fact := 1
		for i := n; i > 0; i-- {
			fact *= i
		}
		out <- fact
		close(out)
	}()
	return out
}
