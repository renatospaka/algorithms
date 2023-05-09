package main

import (
	"fmt"
	"math/rand"
)

func populateValues(size int, random bool) [][]int {
	m := make([][]int, size)
	n := 1;

	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			if random {
				m[r] = append(m[r], rand.Intn(10)-rand.Intn(9))
			} else {
				m[r] = append(m[r], n)
			}	
			n++		
		}
	}
	return m
}

func processSpiralMatrix(matrix [][]int) []int {
	size := len(matrix)
	newM := make([]int, 0)
	left, right := 0, size
	top, bottom := 0, size
	// fmt.Printf("left: %d right: %d top: %d bottom: %d\n", left, right, top, bottom)
	// fmt.Println("==========================================")
	for left < right && top < bottom {
		// get every p in the top row
		for p := left; p < right; p++ {
			newM = append(newM, matrix[top][p])
		}
		top++
		// fmt.Printf("left: %d right: %d top: %d bottom: %d\n", left, right, top, bottom)

		// get every p in the right col
		for p := top; p < bottom; p++ {
			newM = append(newM, matrix[p][right - 1])
		}
		right--
		// fmt.Printf("left: %d right: %d top: %d bottom: %d\n", left, right, top, bottom)

		// get every p in the bottom row
		for p := right - 1; p > left; p-- {
			newM = append(newM, matrix[bottom - 1][p])
		}
		bottom--
		// fmt.Printf("left: %d right: %d top: %d bottom: %d\n", left, right, top, bottom)

		// get every p in the left col
		for p := bottom; p >= top; p-- {
			newM = append(newM, matrix[p][left])
		}
		left++
		// fmt.Printf("left: %d right: %d top: %d bottom: %d\n", left, right, top, bottom)
	}
	return newM
}


func main() {
	size := 5

	m1 := populateValues(size, false)
	fmt.Println("Matrix 1:", m1)

	newMatrix := processSpiralMatrix(m1)
	fmt.Printf("New Matrix: %v Size: %d\n", newMatrix, len(newMatrix))
}
