// Like a C for
// for init; condition; post { }

// Like a C while
// for condition { }

// Like a C for(;;)
// for { }

package main

import "fmt"

func main() {
	//==== Nesting Loops ====
	// for i := 0; i <= 5; i++ {
	// 	fmt.Printf("The outer loop: %d\n ", i)
	// 	for j := 0; j < 3; j++ {
	// 		fmt.Printf("\t\tThe inner loop: %d\n", j)
	// 	}
	// }

	// ==== Other Form ====
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
}
