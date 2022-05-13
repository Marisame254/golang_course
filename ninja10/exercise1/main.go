// get this code working: https://play.golang.org/p/j-EA6003P0
// 		using func literal, aka, anonymous self-executing func
//		a buffered channel

package main

import "fmt"

func main() {
	fmt.Println("Hi :3")

	c := make(chan int, 1)
	// using a func literal
	// go func() {
	// 	c <- 42
	// }()
	// =================

	// buffered solution
	c <- 42
	// ================

	fmt.Println(<-c)
}
