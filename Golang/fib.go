package main

import (
	"fmt"
)

func fib(num int) int {
	if (num == 0 || num == 1) {
		return 1
	}
	return fib(num - 2) + fib(num - 1)
}

func phi(num int) float64 {
	return (float64(fib(num+1)) / float64(fib(num)))
}

func main() {
	fmt.Println("Hello", phi(50))
}
