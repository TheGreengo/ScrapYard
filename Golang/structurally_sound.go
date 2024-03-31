package main

import (
	"fmt"
	"math"
)

type euclid interface {
	distance() float64
}

type coord_2 struct {
	x int
	y int
}

type coord_3 struct {
	x int
	y int
	z int
}

func (c coord_2) distance() float64 {
	return math.Sqrt(float64(c.x * c.x + c.y * c.y))
}

// func (c coord_3) distance() float64 {
// 	return math.Sqrt(float64(c.x * c.x + c.y * c.y + c.z * c.z))
// }

func main() {
	var thing coord_2
	thing.x = 10
	thing.y = -10

	thang := coord_2{ 2, -2 }
	theng := coord_3{ 2, -2, 3 }

	var c2 interface{} = coord_2{}
	var c3 interface{} = coord_3{}

	if _, ok := c2.(euclid); ok {
		fmt.Println("coord_2 implements the distance interface")
	} else {
		fmt.Println("coord_2 does not implement the distance interface")
	}
	
	if _, ok := c3.(euclid); ok {
		fmt.Println("coord_3 implements the distance interface")
	} else {
		fmt.Println("coord_3 does not implement the distance interface")
	}

	fmt.Println(thing)
	fmt.Println(thang)
	fmt.Println(theng)
	fmt.Println(thing.distance())
	fmt.Println(thang.distance())
	//fmt.Println(theng.distance())
}