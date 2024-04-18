package main

import "fmt"

func serious() {
	panic("hay problemas")
}

func main() {
	defer func() {
		if thing := recover(); thing != nil{
			fmt.Println("Your problem is:", thing)
		}
	}()
	serious()
}
