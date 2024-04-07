package main

import (
	"fmt"
	"time"
)

func main() {
	thing := make(chan int, 4)
	tid := time.NewTimer(time.Second)
	tid2 := time.NewTimer(time.Second * time.Duration(3))
	
	for i := 4; i > 0; i -- {
		thing <- i
	}

	close(thing)

	for j := range thing {
		fmt.Println(j)
	}

	<-tid.C

	go func() {
		<-tid2.C
	} ()
	
	thh := tid2.Stop()

	fmt.Println(thh)
}