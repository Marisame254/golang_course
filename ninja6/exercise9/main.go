// A “callback” is when we pass a func into a func as an argument. For this exercise,
// 			● pass a func into a func as an argument

package main

import "fmt"

func main() {
	fmt.Println("Hi")

	g := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		}
		if len(xi) == 1 {
			return 1
		}
		return xi[0] + xi[len(xi)-1]
	}

	x := foo(g, []int{1, 2, 3, 4, 5})
	fmt.Println(x)
}

func foo(f func(xi []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}
