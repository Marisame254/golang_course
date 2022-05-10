package main

import "fmt"

func main() {
	x := 1
	for {
		x++

		if x > 20 {
			break
		}

		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}
}
