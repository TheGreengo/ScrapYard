package main

import "fmt"

type magnitude struct {
	size float64
}

func anyPrint[T any](thing T) {
	fmt.Println(thing)
}

type node[T any] struct {
	val T 
	left * node[T] 
	right * node[T]
}

type vector struct {
	magnitude
	direction float64
}

func main() {
	thing := vector{
		magnitude: magnitude {
			size: 13.0,
		},
		direction: -90.0,
	}

	fmt.Println(thing)
	anyPrint(4)
	anyPrint("hello")

	thang := node[int]{val: 4, left: nil, right: nil} 
	theng := node[int]{val: 6, left: nil, right: nil} 
	thung := node[int]{val: 5, left: &thang, right: &theng} 
	
	fmt.Println("center:", thung.val)
	fmt.Println("left:", (*thung.left).val)
	fmt.Println("right:", (*thung.right).val)
}