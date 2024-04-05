package main

import (
	"fmt"
    "time"
)

func printNums(num int) {
	// time.Sleep(time.Second)
	for i := 0; i < num; i++ {
		fmt.Println(i)
	}
}

func main() {
	go printNums(30)
	go printNums(30)

	time.Sleep(time.Second) // turns out go 
	func(num int){
		for i := 0; i < num; i++ {
			fmt.Println("num:",i)
		}
	} (20)
}