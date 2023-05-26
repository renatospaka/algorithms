package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var (
		size 						int
		originalMatrix	[][]int
		newMatrix  			[]int
	)
	
	size = 5
	originalMatrix = populateValues(size, false)
	newMatrix = processSpiralMatrix(originalMatrix)
	fmt.Printf("The spiral matrix of %v is %v\n", originalMatrix, newMatrix)

	originalMatrix = [][]int{{1,2,3},{4,5,6},{7,8,9}}
	newMatrix = processSpiralMatrix(originalMatrix)
	fmt.Printf("The spiral matrix of %v is %v\n", originalMatrix, newMatrix)

	originalMatrix = [][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}}
	newMatrix = processSpiralMatrix(originalMatrix)
	fmt.Printf("The spiral matrix of %v is %v\n", originalMatrix, newMatrix)

	originalMatrix = [][]int{{1},{12}}
	newMatrix = processSpiralMatrix(originalMatrix)
	fmt.Printf("The spiral matrix of %v is %v\n", originalMatrix, newMatrix)
}

func processSpiralMatrix(matrix [][]int) []int {
	newM := make([]int, 0)
	left, right := 0, len(matrix[0])
	top, bottom := 0, len(matrix)
	// fmt.Printf("left: %d right: %d top: %d bottom: %d\n", left, right, top, bottom)
	// fmt.Println("==========================================")
	for left < right && top < bottom {
		// get every elem in the top row
		for elem := left; elem < right; elem++ {
			newM = append(newM, matrix[top][elem])
		}
		top++
		// fmt.Printf("left: %d right: %d top: %d bottom: %d\n", left, right, top, bottom)

		// get every elem in the right col
		for elem := top; elem < bottom; elem++ {
			newM = append(newM, matrix[elem][right - 1])
		}
		right--
		// fmt.Printf("left: %d right: %d top: %d bottom: %d\n", left, right, top, bottom)

		// verify if the loop is out of range
		if !(left < right && top < bottom) {
			break
		}

		// get every elem in the bottom row
		for elem := right - 1; elem > left; elem-- {
			newM = append(newM, matrix[bottom - 1][elem])
		}
		bottom--
		// fmt.Printf("left: %d right: %d top: %d bottom: %d\n", left, right, top, bottom)

		// get every elem in the left col
		for elem := bottom; elem >= top; elem-- {
			newM = append(newM, matrix[elem][left])
		}
		left++
		// fmt.Printf("left: %d right: %d top: %d bottom: %d\n", left, right, top, bottom)
	}
	return newM
}

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
