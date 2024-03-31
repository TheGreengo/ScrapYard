package main

import (
	"fmt"
)

func doubler() func() int {
	num := 1
	return func() int {
		num *= 2
		return num
	}
}

func sum(a... int32) int32 {
	var res int32 = 0
	for _, num := range a {
		res += num
	}
	return res
}

func greet() {
	fmt.Println("This worked")
}

func hola() (string, string) {
	return "hello", "world"
}

func main() {
	fmt.Println(sum(1,2,3,4,5,6))
	greet()
	word1, word2 := hola()
	fmt.Println(word1, word2)

	doub := doubler()

	fmt.Println(doub())
	fmt.Println(doub())
	fmt.Println(doub())
	fmt.Println(doub())

	var thing func(k int) int

	thing = func(k int) int {
		if (k == 0 || k == 1) {
			return 1
		} else {
			return k * thing(k-1) 
		}
	}

	fmt.Println(thing(5))

}