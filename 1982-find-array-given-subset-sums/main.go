package main

import (
	"fmt"
)

func main() {
	sums := []int{-3,-2,-1,0,0,1,2,3}
	sum := 3

	answer := findArrayGivenSubsetSums(sum, sums)
	fmt.Printf("Solution array for sum of %v = %d is %v\n", sums, sum, answer)
}

func findArrayGivenSubsetSums(sum int, sums []int) []int {
	answer := []int{}
	sumSoFar := 0

	end := -1
	start := 0
	for ix := 0; ix < len(sums); ix++ {
		sumSoFar = 0
		start = ix
		fmt.Printf("finding sum: %d -> \n", sum)
		ixEnd := ix + sum
		if ixEnd > len(sums) {
			ixEnd = len(sums)
		}
		for ixStart := ix; ixStart < ixEnd; ixStart++ {
			sumSoFar += sums[ixStart]
			fmt.Printf("ix: %d, ixStart: %d, value: %d, sumSoFar: %d\n", ix, ixStart, sums[ixStart], sumSoFar)
			// if sumSoFar > sum {
			// 	fmt.Printf("sumSoFar > sum: %b\n", true)
			// 	fmt.Println("")
			// 	break
			// }
			if sumSoFar == sum &&
			ixStart - ix == sum {
				fmt.Printf("sumSoFar == sum: %v\n", true)
				fmt.Println("")
				end = ixStart
				break
			}	
		}
	}

	fmt.Printf("start: %d, end: %d\n", start, end)
	return answer
}

// func findArrayGivenSubsetSumsWitDP(size int, sums []int) []int {
// 	answer := []int{}
// 	sum := returnSumArray(sums)
// 	if sum % 2 != 0 {
// 		return answer
// 	}

// 	// it will hold all possible sum results from the original sums array
// 	// zero is a valid subset of the array
// 	// dynamicProgramingSet := []int{0}
// 	return answer
// }

// func returnSumArray(sums []int) int {
// 	sum := -1

// 	for i := 0; i < len(sums); i++ {
// 		sum += sums[i]
// 	}

// 	return sum
// }
