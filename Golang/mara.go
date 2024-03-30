package main

import ( 
	"fmt"
)

func main() {
	thing := make(map[string]int32)

	thing["this"] = 4
	thing["that"] = 239
	thing["other"] = 4432

	fmt.Println(thing)

	delete(thing, "this")

	_, cont := thing["this"]
	fmt.Println(cont)
	_, cont = thing["that"]
	fmt.Println(cont)

	thang := map[string]string{
		"hello": "world",
		"hola": "mundo",
		"davs": "verden",
	}

	fmt.Println(thang["davs"])

	other := []int32{5,4,3,2,1}

	for ind, num := range other {
		fmt.Printf("%v: %v\n", ind, num)
	}

	for key, val := range thing {
		fmt.Printf("%v: %v\n", key, val)
	}
}