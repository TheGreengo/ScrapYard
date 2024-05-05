package main

import (
	"fmt"
)

func main() {
	var thing byte
	thing = 255

	thing += 5
	
	fmt.Printf("%d\n", thing)	
	
	thing -= 200

	fmt.Printf("%d\n", thing)	
}
