package main

import "fmt"

func main() {
	fmt.Println("Hi :3")

	// set up the pipeline and consume the output
	for v := range sqrt(sqrt(gen(2, 3))) {
		fmt.Println(v)
	}
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
