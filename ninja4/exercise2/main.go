// Using a COMPOSITE LITERAL:
// 	● create a SLICE of TYPE int
// 	● assign 10 VALUES
// ● Range over the slice and print the values out.
// ● Using format printing
// 	○ print out the TYPE of the slice

package main

import "fmt"

func main() {
	x := []int{30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}
	for i, v := range x {
		fmt.Println("index:", i, "value:", v)
	}

	fmt.Printf("Type: %T\n", x)
}
