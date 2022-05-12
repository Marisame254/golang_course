// Closure is when we have “enclosed” the scope of a variable
// in some code block. For this hands-on exercise,
// create a func which “encloses” the scope of a variable:

package main

import "fmt"

func main() {
	fmt.Println("Hi")

	f := foo()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
