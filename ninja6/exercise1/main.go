// Hands on exercise
// 			○ create a func with the identifier foo that returns an int
// 			○ create a func with the identifier bar that returns an int and a string
// 			○ call both funcs
// 			○ print out their results

package main

import "fmt"

func main() {
	fmt.Println("Hi")

	// call func foo
	f := foo()
	fmt.Println("foo int is:", f)

	// call func bar
	f1, f2 := bar()
	fmt.Println("bar int is:", f1, "\nbar string is:", f2)

}

func foo() int {
	return 10
}

func bar() (int, string) {
	return 100, "Hello"
}
