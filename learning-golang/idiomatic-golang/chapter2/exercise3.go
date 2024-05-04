package main

import (
	"fmt"
)


// Write a program with three variable on named b of type byte, one named smallI of type int32, and one named bigI of type uint64. Assign each variable
// the maximum legal value of its type; then add 1 to each variable. Print out their values.

func main() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	fmt.Println(b + 1)
	fmt.Println(smallI + 1)
	fmt.Println(bigI + 1)

}
