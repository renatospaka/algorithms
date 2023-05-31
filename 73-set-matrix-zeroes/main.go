package main

import "fmt"

func main() {
	var (
		array [][]int
		arrayOrig [][]int
	)

	arrayOrig = [][]int{{1,1,1},{1,0,1},{1,1,1}}
	array = [][]int{{1,1,1},{1,0,1},{1,1,1}}
	setZeroes(array)
	fmt.Printf("Setting zeros for %v => %v\n", arrayOrig, array)

	arrayOrig = [][]int{{0,1,2,0},{3,4,5,2},{1,3,1,5}}
	array = [][]int{{0,1,2,0},{3,4,5,2},{1,3,1,5}}
	setZeroes(array)
	fmt.Printf("Setting zeros for %v => %v\n", arrayOrig, array)

	arrayOrig = [][]int{{3,4,5,2},{0,1,2,0},{1,3,1,5}}
	array = [][]int{{3,4,5,2},{0,1,2,0},{1,3,1,5}}
	setZeroes(array)
	fmt.Printf("Setting zeros for %v => %v\n", arrayOrig, array)

	arrayOrig = [][]int{{3,4,5,2},{0,1,2,0},{3,0,0,5}}
	array = [][]int{{3,4,5,2},{0,1,2,0},{3,0,0,5}}
	setZeroes(array)
	fmt.Printf("Setting zeros for %v => %v\n", arrayOrig, array)
}

func setZeroes(matrix [][]int) {
	// n(1)
	rows, cols, rowZero := len(matrix), len(matrix[0]), false

	// check what rows and cols need to be zero
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if matrix[r][c] == 0 {
				matrix[0][c] = 0
				if r > 0 {
					matrix[r][0] = 0
				} else {
					rowZero = true
				}
			}
		}
	}

	// starting updating rows vs cols to zero, if applys
	for r := 1; r < rows; r++ {
		for c := 1; c < cols; c++ {	
			if matrix[0][c] == 0 || matrix[r][0] == 0 {
				matrix[r][c] = 0
			}
		}
	}

	// do the same for the 1st col
	if matrix[0][0] == 0 {
		for r := 0; r < rows; r++ {
			matrix[r][0] = 0
		}
	}
	
	// do the same for the 1st row
	if rowZero {
		for c := 0; c < cols; c++ {
			matrix[0][c] = 0
		}
	}
}
