package main

// next up, slices

/*
*/

import "fmt"

func main() {
	var nums [3]int
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3

	others := [5]int{8, 7, 6, 5, 4}

	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println(others)
	fmt.Println(len(others))
}