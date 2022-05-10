package main

import "fmt"

func main() {
	// Composite Literal
	x := []int{4, 5, 6, 7, 32}

	fmt.Println(x)
	fmt.Println(x[3:])
	fmt.Println(x[:3])
	fmt.Println(x[1:4])
}

// a SLICE allow yo to group values of the same type
