// Create a value and assign it to a variable.
// Print the address of that value.

package main

import "fmt"

func main() {
	fmt.Println("Hi")

	x := 54
	fmt.Println(x)
	fmt.Println(&x)
}
