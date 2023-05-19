package main

import "fmt"

func main() {
	var (
		numbers []int
		target  int
		result  []int
	)

	numbers = []int{2, 7, 11, 15}
	target = 9
	result = twoSum(numbers, target)
	fmt.Printf("The two numbers indexes of array %v that sum %d is %v\n", numbers, target, result)

	numbers = []int{2, 3, 4}
	target = 6
	result = twoSum(numbers, target)
	fmt.Printf("The two numbers indexes of array %v that sum %d is %v\n", numbers, target, result)

	numbers = []int{-1, 0}
	target = -1
	result = twoSum(numbers, target)
	fmt.Printf("The two numbers indexes of array %v that sum %d is %v\n", numbers, target, result)
}

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}

		if sum > target {
			right--
		} else if sum < target {
			left++
		}
	}

	return []int{}
}
