package main

import (
	"fmt"
)

func ffffft(word chan<- string, ward string) {
	word <- ward
}

func tfffff(word <-chan string) {
	ward := <-word
	fmt.Println(ward)
}

func main() {
	thing := make(chan string, 1)

	ffffft(thing, "butt")
	tfffff(thing)
}