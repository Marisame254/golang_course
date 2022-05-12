package main

import "fmt"

func main() {
	foo()
	bar("James")

	s1 := woo("Maria")
	fmt.Println(s1)

	x, y := mouse("Carlos", "Fleming")
	fmt.Println(x, y)
}

// Basic function
func foo() {
	fmt.Println("Hello from foo")
}

// function with arguments
func bar(s string) {
	fmt.Println("hello", s)
}

// function with return
func woo(st string) string {
	return fmt.Sprintln("Hello from woo, ", st)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, `, says "Hello"`)
	b := false
	return a, b
}
