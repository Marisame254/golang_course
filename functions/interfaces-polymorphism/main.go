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

func (s secretAgent) speak() {
	fmt.Println("Im am", s.first, s.last, "the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("Im am", p.first, p.last, "-  the person speak")
}

type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into bar", h.(person).first)
	case secretAgent:
		fmt.Println("I was passes into bar", h.(secretAgent).first)
	}
	fmt.Println("I was passes into bar", h)
}

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

	p1 := person{
		first: "Dr",
		last:  "Yes",
	}

	fmt.Println(sa1)
	sa1.speak() // use method
	sa2.speak()

	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)
}
