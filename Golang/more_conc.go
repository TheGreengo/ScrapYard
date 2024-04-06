package main

import (
	"fmt"
	"time"
)

func ffffft(word chan<- string, ward string) {
	word <- ward
}

func tfffff(word <-chan string) {
	ward := <-word
	fmt.Println(ward)
}

func timeDel(word chan<- string, del int, mess string) {
	time.Sleep(time.Duration(del) * time.Second)
	word <- mess
}

func main() {
	thing := make(chan string, 1)
	thang := make(chan string, 1)
	theng := make(chan string, 1)

	ffffft(thing, "butt")
	tfffff(thing)

	go timeDel(thang, 1, "first")
	go timeDel(theng, 2, "second")

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-thang:
			fmt.Println(msg1)
		case msg2 := <-theng:
			fmt.Println(msg2)
		}
	}

	go timeDel(thang, 3, "third")
	select {
	case msg1 := <-thang:
		fmt.Println(msg1)
	case <-time.After(time.Second):
		fmt.Println("delayed")
	}
}