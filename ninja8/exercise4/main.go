//Starting with this code, sort the []int and []string for each person.
// THUS LINK: https://play.golang.org/p/H_q75mpmHW

package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	// sorting the ints
	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	// sorting the strings
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
