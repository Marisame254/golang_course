package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// Method ============
func (s secretAgent) speak() {
	fmt.Println("Im am", s.first, s.last)
}

//====================

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Carlos",
			"Trejo",
		},
		ltk: true,
	}

	fmt.Println(sa1)
	sa1.speak() // use method
	sa2.speak()
}
