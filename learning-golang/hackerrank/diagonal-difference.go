package main

import (
    "fmt"
)

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	var leftToRightSum, rightToLeftSum int32

	if len(arr) > 0 { // check if matrix is empty
		for i := 0; i < len(arr); i++ {
			leftToRightSum += arr[i][i] // add elements on the diagonal from left to right
			rightToLeftSum += arr[len(arr)-1-i][i] // add elements on the diagonal from right to left
		}
		return (leftToRightSum - rightToLeftSum)
	}

	return 0
}


func main() {
	
	var matrix [][][]int32

	depth := 3
	rows := 3
	cols := 3

	matrix = make([][][]int32, depth)
	for i := range matrix {
		matrix[i] = make([][]int32, rows)
		for j := range matrix[i] {
			matrix[i][j] = make([]int32, cols)
		}
	}

	fmt.Println(matrix)
	fmt.Println("hello there")

	// result := diagonalDifference(arr)
}

