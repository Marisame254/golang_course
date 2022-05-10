// Using iota, create 4 constants for the
// NEXT 4 years. Print the constant values.

package main

import "fmt"

const (
	a = 2022 + iota
	b = 2022 + iota
	c = 2022 + iota
	d = 2022 + iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
