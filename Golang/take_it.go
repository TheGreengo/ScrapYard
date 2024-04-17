package main

import (
	"fmt"
	"time"
)

func main() {
	thing := make(chan int, 4)

	for i := 0; i < 4; i++ {
		thing <- (i * i)
	}
	close(thing)

	limit := time.Tick(time.Duration(200) * time.Millisecond)

	for thang := range thing {
		<- limit
		fmt.Println("Thing:", thang, "time:", time.Now())
	}
	    
}
