// Fix the race condition you created in
// exercise #4 by using package atomic

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("Hi :3")

	var incrementer int64
	var wg sync.WaitGroup

	gs := 50
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			fmt.Println(atomic.LoadInt64(&incrementer))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("end value:", incrementer)
}
