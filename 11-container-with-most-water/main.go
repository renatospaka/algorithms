package main

import "fmt"

func main() {
	test([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	test([]int{1, 1})
	test([]int{8, 7, 2, 1})
}

func test(nums []int) {
	area := maxArea(nums)
	fmt.Printf("The max area of %v is %v\n", nums, area)
}

func maxArea(height []int) int {
	result, left, right := 0, 0, len(height)-1

	for left < right {
		area := (right - left) * min(height[left], height[right])
		result = max(result, area)

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
