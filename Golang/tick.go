package main 

import (
	"fmt"
	"time"
)

func main() {
	tiktok := time.NewTicker(time.Millisecond * time.Duration(50))
	thing := make(chan bool)

	go func() {
		for {
			select {
			case <-thing:
				fmt.Println("number")
				return
			case t:= <-tiktok.C:
				fmt.Println("@:",t)
			}
		}
	} ()

	time.Sleep(time.Second)
	tiktok.Stop()
	thing <- true
	fmt.Println("done")
}