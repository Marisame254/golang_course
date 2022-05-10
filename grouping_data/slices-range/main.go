package main

import "fmt"

func main() {
	// Composite Literal
	x := []int{4, 5, 6, 7, 32}

	for i, v := range x {
		fmt.Println(i, v)
	}
}

// a SLICE allow yo to group values of the same type
