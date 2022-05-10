package main

import "fmt"

func main() {
	m := map[string]int{
		"James":    34,
		"Jenkings": 21,
		"Maria":    30,
	}
	fmt.Println(m)

	delete(m, "James")
	fmt.Println(m)

	delete(m, "Maria")
	fmt.Println(m)

	if v, ok := m["Jenkings"]; ok {
		fmt.Println("value", v)
		delete(m, "Jenkings")
	}

	fmt.Println(m)
}
