package main 

import (
	"fmt"
)

func main() {
	num := 0
	spot := &num

	var nam int
	var spat *int

	spat = &nam
	*spat = 4
	fmt.Println(num)
	fmt.Println(nam)
	num = 5
	fmt.Println(*spot)
}