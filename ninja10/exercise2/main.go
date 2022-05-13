// Get this code running:
// https://play.golang.org/p/oB-p3KMiH6

package main

import "fmt"

func main() {
	fmt.Println("Hi :3")

	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
