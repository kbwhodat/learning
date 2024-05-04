package main

import (
	"fmt"
)

// Write a program that declares a constant called value that can be assigned to both an integer and a floating-point variable. Assign it to an integer
// called i and a floating-point variable called f. Print out i and f.

func main() {
	const Value = 10

	var i int = Value
	var f float32 = Value

	fmt.Println(i)
	fmt.Println(f)

}
