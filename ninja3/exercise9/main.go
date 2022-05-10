// Create a program that uses a switch statement
// with the switch expression specified as a
// variable of TYPE string with the IDENTIFIER “banana”.

package main

import "fmt"

func main() {
	x := "banana"
	switch x {
	case "banana":
		fmt.Println("It's a banana")
	case "Melon":
		fmt.Println("It's a Melon")
	case "pineapple":
		fmt.Println("It's a pineapple")
	case "apple":
		fmt.Println("It's a apple")
	}
}
