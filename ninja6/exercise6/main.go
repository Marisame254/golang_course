// Build and use an anonymous func

package main

import "fmt"

func main() {
	fmt.Println("Hi")

	func() {
		for i := 0; i < 50; i++ {
			fmt.Println(i)
		}
	}()
}
