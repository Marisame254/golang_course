package main

import "fmt"

func main() {
	// anonymous struct
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   35,
	}

	fmt.Println(p1)

	fmt.Println(p1.first, p1.last, p1.age)
}
