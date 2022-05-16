package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hi :3")

	// first example ==========
	// n, err := fmt.Println("Hello")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(n)
	//=========================

	// second example ===============
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("James bond")

	io.Copy(f, r)
	//===============================
}
