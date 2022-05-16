package main

import (
	"fmt"

	"github.com/Marisame254/golang_course/ninja12/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fmt.Println("Hi :3")

	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
}
