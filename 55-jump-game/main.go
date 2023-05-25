package main

import "fmt"

func main() {
	var (
		numbers    []int
		reacheable bool
	)

	numbers = []int{2, 3, 1, 1, 4}
	reacheable = canJump(numbers)
	fmt.Printf("Is it possible to go throught %v till the end? %t\n", numbers, reacheable)
}

// using greed algorithm
func canJump(nums []int) bool {
	goal := len(nums) - 1		// coming backwards

	for i := len(nums) - 1; i >= 0; i-- {
		if i + nums[i] >= goal {
			goal = i
		}
	}
	return goal == 0
}
