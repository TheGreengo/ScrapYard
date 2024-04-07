package main

import (
	"fmt"
)

func main() {
	thing := make(chan int, 4)
	
	for i := 4; i > 0; i -- {
		thing <- i
	}

	close(thing)

	for j := range thing {
		fmt.Println(j)
	}
}