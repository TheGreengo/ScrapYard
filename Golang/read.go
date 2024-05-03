package main

import (
	"fmt"
	"os"
)

/*
* ok, so I want to make a little nifty encryption program
* I want it to:
* - take in command line arguments (file name, encrypt/decrypt, key)
* - then I want it to read every byte, and change it to being the 
*/

func main() {
	dat, err := os.ReadFile("./test.txt")
	
	if err != nil {
		panic(err)
	}

	fmt.Println(len(dat))	
}
