package main

import (
	"fmt"
	"time"
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

	theng := make(chan int, 4)

	go func () {
		for {
			i, more := <-theng
			if !more {
				fmt.Println("donzo")
				return
			}
			fmt.Println(i)
		}
	} ()

	for j := 0; j < 4; j++ {
		theng <- j
	}

	time.Sleep(time.Second)

	close(theng)
	_, good := <-theng
	if (good) {
		fmt.Println("Closed successfully")
	}
}