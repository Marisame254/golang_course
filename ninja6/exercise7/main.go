// Assign a func to a variable, then call that func

package main

import "fmt"

func main() {
	fmt.Println("Hi")

	f := func() {
		for i := 0; i < 20; i++ {
			fmt.Println(i)
		}
	}

	f()
}
