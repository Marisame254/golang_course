package main

import "fmt"

func main() {
	fmt.Println("Hi :3")

	// Set up the pipeline
	c := gen(2, 3)
	out := sqrt(c)

	// Consume the outputs
	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9
}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums { // need the _, n
			out <- n
		}
		close(out)
	}()
	return out
}

func sqrt(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
