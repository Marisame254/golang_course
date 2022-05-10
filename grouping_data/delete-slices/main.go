package main

import "fmt"

func main() {
	// Composite Literal
	x := []int{4, 5, 6, 7, 8}
	fmt.Println(x)
	x = append(x, 77, 78, 79, 80)
	fmt.Println(x)

	y := []int{100, 200, 300, 400, 500}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:2], x[6:]...)
	fmt.Println(x)
}

// a SLICE allow yo to group values of the same type
