package main

import "fmt"

type hotdog int

var x hotdog

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 45
	fmt.Println(x)
	y := int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}
