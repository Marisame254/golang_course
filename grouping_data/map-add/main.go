package main

import "fmt"

func main() {
	m := map[string]int{
		"James":    32,
		"Federico": 27,
	}

	fmt.Println(m)
	fmt.Println(m["James"])

	m["Maria"] = 30

	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, v)
	}
}
