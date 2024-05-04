package main

import (
	"os"
)

func main() {
	writ_file, _ := os.Create("./written.txt")
	defer writ_file.Close()

	mess := []byte{ 66, 114, 97, 100, 101, 110, 10 }			
	writ_file.Write(mess)
}
