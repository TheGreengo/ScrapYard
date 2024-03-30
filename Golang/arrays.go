package main

// next up, slices

/*
*/

import (
	"fmt"
	"slices"
)

func main() {
	var nums [3]int32
	var slice []float32

	slice = make([]float32, 5, 5)
	
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3

	for i := 0; i < 5; i++ {
		slice[i] = float32(i * i)
	}

	var ecils []float32
	ecils = make([]float32, len(slice))
	copy(ecils, slice)

	if (slices.Equal(ecils, slice)) {
		fmt.Println("They match")
	} else {
		fmt.Println("They don't")
		fmt.Println(ecils)
	}

	slice = append(slice, 4.0)
	slice = append(slice, 6.0005)
	slice = append(slice, -7.000404020)


	if (slices.Equal(ecils, slice)) {
		fmt.Println("They match")
	} else {
		fmt.Println("They don't")
	}

	others := [5]int32{8, 7, 6, 5, 4}

	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println(others)
	fmt.Println(len(others))

	fmt.Println(slice)
	fmt.Println(len(slice))
}