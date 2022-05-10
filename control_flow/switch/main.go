package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print2")
	case (3 == 3):
		fmt.Println("prints")
		fallthrough
	case (4 == 4):
		fmt.Println("also true, does it prints?")
		fallthrough
	case (11 == 14):
		fmt.Println("not true 7u7")
		fallthrough
	case (15 == 15):
		fmt.Println("true *0*")
		fallthrough
	default:
		fmt.Println("this is default")
	}
}
