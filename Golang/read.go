package main

import (
	"fmt"
	"os"
	"bufio"
	"errors"
	"io"
)

/*
* ok, so I want to make a little nifty encryption program
* I want it to:
* - take in command line arguments (file name, encrypt/decrypt, key)
* - then I want it to read every byte, and change it to being the 
* here we're just going to open a file and then read the bits
*/

func main() {
	file, err := os.Open("./test.txt")
	
	if err != nil {
		panic(err)
	}
	
	in_the_buff := bufio.NewReader(file)
	
	for {
		b, er := in_the_buff.ReadByte()	
	
		if errors.Is(er, io.EOF) {
			fmt.Printf("%d\n", b)
			fmt.Println("End of line")
			break
		}

		if er != nil {
			panic(er)
		}
	
		fmt.Printf("%d\n", b)
	}
}
