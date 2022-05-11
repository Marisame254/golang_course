// Create a slice of a slice of string ([][]string).
// Store the following data in the multi-dimensional slice:
// "James", "Bond", "Shaken, not stirred"
// "Miss", "Moneypenny", "Helloooooo, James."
// Range over the records, then range over the data in each record.

package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James"}

	fmt.Println(x)
	fmt.Println(y)

	xy := [][]string{x, y}

	for i, slice := range xy {
		fmt.Println("Index:", i)
		for j, name := range slice {
			fmt.Println("\tindex:", j, "values:", name)
		}
	}
}
