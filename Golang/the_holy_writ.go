package main

import (
	"os"
	"fmt"
)

func main() {
	
	mess := []byte{ 66, 114, 97, 100, 101, 110, 10 }			

	 err := os.WriteFile("./written.txt", mess, 0644)

	if err != nil {
		fmt.Println("didn't open")
		panic(err)
	}
}
