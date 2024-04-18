package main

import (
    "fmt"
	"os"
)

func main() {
    fmt.Printf("normal thing %v\n", 64)
	fmt.Printf("binary: %b\n", 14)
	fmt.Printf("hex: %x\n", 20321753)
    fmt.Printf("these should be space %6d %6.2f and then %-6d %-6.2f\n", 2, 3.1415, 420, 69.0)
    fmt.Printf("can you booly stuff? %t %t %t\n", true, false, true)
    fmt.Fprintf(os.Stderr, "this is an error\n")
}
