package main

import "fmt"

func main() {
	test([]int{2, 7, 11, 15}, 9)
	test([]int{2, 3, 4}, 6)
	test([]int{-1, 0}, -1)
}

func test(numbers []int, target int) {
	result := twoSum(numbers, target)
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
