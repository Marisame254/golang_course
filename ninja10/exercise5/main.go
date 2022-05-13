//Show the comma ok idiom starting with this code.
// THIS CODE: https://play.golang.org/p/YHOMV9NYKK

package main

import "fmt"

func main() {
	fmt.Println("Hi :3")

	c := make(chan int)

	go func() {
		c <- 42
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)

}
