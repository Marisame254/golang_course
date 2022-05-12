package main

import "fmt"

func main() {
	fmt.Println("hi")

	f := func() {
		fmt.Println("my first func expression")
	}
	f()

	g := func(x int) {
		fmt.Println("The number is:", x)
	}
	g(32)

}
