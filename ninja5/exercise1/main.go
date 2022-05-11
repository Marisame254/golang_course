// Create your own type “person” which will have an underlying type of “struct” so that it can store
// the following data:
// ● first name
// ● last name
// ● favorite ice cream flavors
// Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice
// which stores the favorite flavors.

package main

import "fmt"

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {

	p1 := person{
		first:      "James",
		last:       "Bond",
		favFlavors: []string{"chocolate", "coke"},
	}

	p2 := person{
		first:      "Nue",
		last:       "Chen",
		favFlavors: []string{"vanilla", "strawberry"},
	}

	fmt.Println(p1.first, p1.last, p1.favFlavors[0], p1.favFlavors[1])
	fmt.Println(p2.first, p2.last, p2.favFlavors[0], p2.favFlavors[1])
}
