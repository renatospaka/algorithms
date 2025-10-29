package main

import "fmt"

func main() {
	test([]int{1, 1, 1}, 2)
	test([]int{1, 2, 3}, 3)
	// test([]int{1, 1, 1, 1, 1, 1}, 3)
	// test([]int{1, -1, 1, 1, 1, 1}, 3)
}

func test(nums []int, target int) {
	result := subArraySum(nums, target)
	fmt.Printf("The array %v has %d subarrays that sum to %d\n", nums, result, target)
}

func subArraySum(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	prefixSum := make(map[int]int)
	prefixSum[0] = 1
	currentSum, sum := 0, 0
	for i := 0; i < len(nums); i++ {
		currentSum += nums[i]
		diff := currentSum - target

		if d, ok := prefixSum[diff]; ok {
			sum += d
		}

		if d, ok := prefixSum[currentSum]; ok {
			prefixSum[currentSum] = 1 + d
			continue
		}

		prefixSum[currentSum] = 1
	}
	return sum
}
