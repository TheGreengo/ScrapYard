package main

import (
	"fmt"
)

func main() {
	thing := make(chan bool)

	select {
	case thang := <-thing:
		fmt.Println(thang)
	default:
		fmt.Println("can you have")
		fmt.Println("two lines")
	}

	thang := true
	select {
	case thing <- thang:
		fmt.Println("sent")
	default:
		fmt.Println("Why don't they just use more {s")
	}

	go func() {thing <- true}()
	val := <-thing
	fmt.Println(val)
}