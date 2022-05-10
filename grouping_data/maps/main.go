package main

import "fmt"

func main() {
	m := map[string]int{
		"James":    32,
		"Federico": 27,
	}

	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Maria"])

	if v, ok := m["James"]; ok {
		fmt.Println("This is the if print", v)
	}
}
