package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Println(s1)

	// take function
	s2 := bar()
	fmt.Printf("%T\n", s2)

	// use the function
	x := s2()
	fmt.Println(x)

	fmt.Println(s2())    //using just s2()
	fmt.Println(bar()()) // using just bar()()
}

// normal function
func foo() string {
	return "Hello world"
}

// Return a function
func bar() func() int {
	return func() int {
		return 32
	}
}
