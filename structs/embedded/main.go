package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   35,
		},
		ltk: true,
	}

	fmt.Println(sa1)

	fmt.Println(sa1.person.first, sa1.person.last, sa1.person.age, sa1.ltk)

}
