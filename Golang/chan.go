package main

import (
	"fmt"
	"time"
)

func blocker(start chan bool) {
	time.Sleep(time.Second)
	start <- false
}

func main() {
	finished := make(chan bool, 1)
	thing := make(chan string)
	mess := make(chan string, 3)

	mess <- "this"
	mess <- "that"
	mess <- "the other"

	go func(){
		thing <- "hello"
	}()

	thang := <- thing
	fmt.Println(thang)

	mass := <- mess
	moss := <- mess
	muss := <- mess

	fmt.Println(mass, moss, muss)

	fmt.Println("Launching the stopper")

	go blocker(finished)

	<- finished

	fmt.Println("Waited")
}