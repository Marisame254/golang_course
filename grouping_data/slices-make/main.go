package main

import "fmt"

func main() {
	// Composite Literal
	x := make([]int, 10, 100)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 32
	x[9] = 999

	x = append(x, 123)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 456)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 789)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}

// a SLICE allow yo to group values of the same type
