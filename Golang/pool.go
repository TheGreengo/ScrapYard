package main

import (
	"fmt"
	"time"
)

func thing_doer(name string, thing chan int, out chan int) {
	for i := range thing {
		fmt.Println("Running job:",i ,"on:", name)
		time.Sleep(time.Second)
		fmt.Println("Done with job:",i ,"on:", name)
		out <- i
	}
}

func main() {
	things := make(chan int, 8)
	outs := make(chan int, 8)

	names := []string{"a","b","c","d"}

	for i := 0; i < 4; i++ {
		go thing_doer(names[i], things, outs)
	}
	
	for i := 0; i < 9; i++ {
		things <- i
	}
	
	for i := 0; i < 9; i++ {
		num := <-outs
		fmt.Println(num)
	}
	
}