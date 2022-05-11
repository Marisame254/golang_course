// Using the code from the previous example, add a record
// to your map. Now print the map out
// using the “range” loop

package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	//fmt.Println(m)

	m["Juanita"] = []string{"games", "skateboarding", "pizza"}

	for k, v := range m {
		fmt.Println("Key:", k)
		for i, name := range v {
			fmt.Println("\tindex:", i, "value:", name)
		}
	}
}
