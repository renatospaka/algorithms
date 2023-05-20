package main

import "fmt"

func main() {
	var (
		numbers []int
		target  int
		result  int
	)

	numbers = []int{4,5,6,7,0,1,2}
	target = 0
	result = search(numbers, target)
	fmt.Printf("The index of %v for target = %d is %d\n", numbers, target, result)

	numbers = []int{4,5,6,7,0,1,2}
	target = 3
	result = search(numbers, target)
	fmt.Printf("The index of %v for target = %d is %d\n", numbers, target, result)

	numbers = []int{1}
	target = 0
	result = search(numbers, target)
	fmt.Printf("The index of %v for target = %d is %d\n", numbers, target, result)

	numbers = []int{3,4,5,6,7,0,1,2}
	target = 6
	result = search(numbers, target)
	fmt.Printf("The index of %v for target = %d is %d\n", numbers, target, result)
}

func search(nums []int, target int) int {
	left, right, middle := 0, len(nums) - 1, nums[0]

	for left <= right {
		middle = (left + right) / 2
		if nums[middle] == target {
			return middle
		}

		if nums[left] <= nums[middle] {
			if target > nums[middle] || target < nums[left] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		} else {
			if target < nums[middle] || target > nums[right] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
	}

	return -1
}
