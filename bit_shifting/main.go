package main

import "fmt"

//==== Third Example ====
const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	//==== First Example =====
	//  x := 8
	// 	fmt.Printf("%d\t\t%b\n", x, x)

	// 	y := x << 1
	// 	fmt.Printf("%d\t\t%b\n", y, y)

	//==== Second Example ====
	// kb := 1024
	// mb := kb * 1024
	// gb := mb * 1024

	// fmt.Printf("%d\t\t%b\n", kb, kb)
	// fmt.Printf("%d\t\t%b\n", mb, mb)
	// fmt.Printf("%d\t%b\n", gb, gb)

	//==== Third Example ====

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)
}
