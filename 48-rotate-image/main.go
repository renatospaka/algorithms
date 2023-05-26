package main

import "fmt"

func main() {
	var (
		matrix   [][]int
		original [][]int
	)

	original = [][]int{{1,2,3},{4,5,6},{7,8,9}}
	matrix = [][]int{{1,2,3},{4,5,6},{7,8,9}}
	rotate(matrix)
	fmt.Printf("The 90º rotated matrix of %v is => %v\n", original, matrix)

	original = [][]int{{5,1,9,11},{2,4,8,10},{13,3,6,7},{15,14,12,16}}
	matrix = [][]int{{5,1,9,11},{2,4,8,10},{13,3,6,7},{15,14,12,16}}
	rotate(matrix)
	fmt.Printf("The 90º rotated matrix of %v is => %v\n", original, matrix)
}

func rotate(matrix [][]int) {
	left, right, top, bottom := 0, len(matrix) - 1, 0, 0

	for left < right {
		// fmt.Printf("left: %d || right: %d || right - left : %d\n", left, right, right - left)
		for i := 0; i < (right - left); i++ {
			// fmt.Printf("i: %d || left: %d || right: %d\n", i, left, right)
			top, bottom = left, right

			// save the top left corner value (initial sopt)
			topLeft := matrix[top][left + i]

			// move bottom/left to to top/left
			matrix[top][left + i] = matrix[bottom - i][left]

			// move bottom/right to bottom/left
			matrix[bottom - i][left] = matrix[bottom][right - i]

			// move top/right to bottom/right
			matrix[bottom][right - i] = matrix[top + i][right]

			// move top/left to top/right
			matrix[top + i][right] = topLeft
		}
		right--
		left++
	}
}
