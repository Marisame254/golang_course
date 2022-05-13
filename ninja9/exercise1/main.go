// in addition to the main goroutine, launch two additional goroutines
//		 ○ each additional goroutine should print something out
// use waitgroups to make sure each goroutine finishes before your program exists

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	fmt.Println("Begin CPUs:", runtime.NumCPU())
	fmt.Println("Begin Goroutines:", runtime.NumGoroutine())

	wg.Add(2)
	go func() {
		fmt.Println("Hi :3")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("about to exit")
	fmt.Println("End CPUs:", runtime.NumCPU())
	fmt.Println("End Goroutines:", runtime.NumGoroutine())
}
