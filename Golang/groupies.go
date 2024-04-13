package main

import (
	"fmt"
	"time"
	"sync"
)

func work(id int) {
	fmt.Println("running", id)
	time.Sleep(time.Second)
	fmt.Println(id, "run")
}

func main() {
	var thing sync.WaitGroup

	for i := 0; i < 5; i++ {
		thing.Add(1)

		go func() {
			defer thing.Done()
			work(i)
		} ()
	}

	thing.Wait()
}