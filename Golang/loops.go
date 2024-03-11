package main

import "fmt"

func main() {
	i := 0 
	for i < 3 {
		fmt.Println(i)
		i++
	}

	for j := 0; j < 4; j++ {
		if j % 2 == 0 {
			fmt.Printf("%d is even\n", j)
		} else if j % 3 == 0 {
			fmt.Printf("%d is divisible by three\n", j)
		}
	}

	for j := 0; j < 4; j++ {
		switch j {
		case 3:
			fmt.Println("number three")
		case 0, 2:
			fmt.Println("even") 
		}
	}
}