package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var hydro atomic.Uint64
	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(1)
		
		go func() {
			for j := 0; j < 10000; j++ {
				hydro.Add(1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(hydro.Load())
}
