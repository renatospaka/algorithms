package main

import "fmt"

func main() {
	var (
		numbers []int
		minimum int
	)

	numbers = []int{3, 4, 5, 1, 2}
	minimum = findMin(numbers)
	fmt.Printf("The original array is %v andits minimum is : %d\n", numbers, minimum)

	numbers = []int{4, 5, 6, 7, 0, 1, 2}
	minimum = findMin(numbers)
	fmt.Printf("The original array is %v andits minimum is : %d\n", numbers, minimum)

	numbers = []int{11, 13, 15, 17}
	minimum = findMin(numbers)
	fmt.Printf("The original array is %v andits minimum is : %d\n", numbers, minimum)

	numbers = []int{0, 1, 2, 3, 4, 5}
	minimum = findMin(numbers)
	fmt.Printf("The original array is %v andits minimum is : %d\n", numbers, minimum)
}

func findMin(nums []int) int {
	left, right, middle, minimum := 0, len(nums)-1, 0, nums[0]
	for left < right {
		middle = (left + right) / 2
		minimum = min(minimum, nums[middle])
		// fmt.Printf("middle: %d | minimum: %d | nums[middle]: %d\n", middle, minimum, nums[middle])
		// fmt.Printf("left: %d | nums[left]: %d | right: %d | nums[right]: %d\n", left, nums[left], right, nums[right])

		if nums[middle] > nums[right] {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return min(minimum, nums[left])
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
