// create a type SQUARE
// create a type CIRCLE
// attach a method to each that calculates AREA and returns it
// 			○ circle area= π r 2
// 			○ square area = L * L
// create a type SHAPE that defines an interface as anything that has the AREA method
// create a func INFO which takes type shape and then prints the area
// create a value of type circle
// use func info to print the area of square
// use func info to print the area of circle

package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	lenght float64
}

func (s square) area() float64 {
	return s.lenght * s.lenght
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("the area is:", s.area())
}

func main() {
	fmt.Println("Hi")

	c := circle{
		radius: 12.123,
	}

	s := square{
		lenght: 15,
	}

	info(c)
	info(s)
}
