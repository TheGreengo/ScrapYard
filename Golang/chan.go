package main

import (
	"fmt"
)

func main() {
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
}