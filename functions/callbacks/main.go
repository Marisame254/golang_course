package main

import "fmt"

func main() {

	// Call to sum function
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(x)

	s := sum(x...)
	fmt.Println("Sum numbers", s)

	// Call to even function (Callback)
	s2 := even(sum, x...)
	fmt.Println("Sum even numbers:", s2)

	// Call the odd function (Callback)
	s3 := odd(sum, x...)
	fmt.Println("Sum odd numbers:", s3)

}

// Sum  function
func sum(x1 ...int) int {
	//fmt.Printf("%T\n", x1)
	total := 0
	for _, v := range x1 {
		total += v
	}
	return total
}

// Even function
func even(f func(x1 ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

// Odd function
func odd(f func(x1 ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
