package main

import (
	"fmt"
)

// Write a program that declares an integer variable called i with the value 20.
// Assign i to a floating-point variable name f. Print out i and f.
func main() {
	var i int = 20
	var f float32 = float32(i)

	fmt.Println(i)
	fmt.Println(f)
}
