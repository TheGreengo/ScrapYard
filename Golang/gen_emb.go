package main

import "fmt"

type magnitude struct {
	size float64
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
}