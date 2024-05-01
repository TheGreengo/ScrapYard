package main

import (
	"fmt"
)

func fib(num int) int {
	if (num == 1 || num == 2) {
		return 1
	}
	return fib(num - 2) + fib(num - 1)
}

func phi(num int) float64 {
	return (float64(fib(num+1)) / float64(fib(num)))
}

func main() {
	var i int
	fmt.Scan(&i)
	fmt.Println("Fib",i,":", fib(i))
	fmt.Println("Phi",i,":", phi(i))
}
