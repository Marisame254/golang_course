// Fix the race condition you created in the previous exercise by using a mutex
// 			● it makes sense to remove runtime.Gosched()

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hi :3")

	incrementer := 0

	var wg sync.WaitGroup
	var m sync.Mutex
	gs := 20
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			v := incrementer
			v++
			incrementer = v
			fmt.Println(incrementer)
			m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("end value:", incrementer)
}
