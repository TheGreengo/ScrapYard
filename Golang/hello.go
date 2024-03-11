package main

import (
    "fmt"
    "math"
)

const pi = 3.14159265358

func main() {
	thing := "Hello"
	var thang int = 4 + 6
	var a, b, c int = 1, 2, 3
	num, nam, nom := 5, 6, 7
	val := true && false
	var mask uint16 = 0b1111111100000000 & 0x0ff0 
	
	fmt.Println(thing)
	fmt.Println(thang)
	fmt.Println(a, b, c)
	fmt.Println(num, nam, nom)
	fmt.Println(val)
	fmt.Println(mask)
	fmt.Println(math.Sin(pi))
}