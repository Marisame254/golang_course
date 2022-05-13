// write a program that
//		○ launches 10 goroutines
// 		■ each goroutine adds 10 numbers to a channel
// 		○ pull the numbers off the channel and print them

package main

import "fmt"

func main() {
	fmt.Println("Hi :3")

	c := make(chan int)

	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}

	fmt.Println("about to exit")

}
