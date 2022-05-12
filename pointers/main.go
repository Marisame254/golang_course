package main

import "fmt"

func main() {
	fmt.Println("Hi")

	a := 32
	fmt.Println(a)
	fmt.Println(&a) // & show the address in memory

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a) //type pointer

	b := &a
	fmt.Println(b)  // same address (& show address)
	fmt.Println(*b) // pointing to an address (* show the value in the address)
	fmt.Println(*&a)

	*b = 30
	fmt.Println(a)
}
