package main

import "fmt"

func main() {
	var (
		numbers []int
		sum     int
	)

	numbers = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	sum = maxSubArray(numbers)
	fmt.Printf("The array %v has the largest sum = %d\n", numbers, sum)

	numbers = []int{1}
	sum = maxSubArray(numbers)
	fmt.Printf("The array %v has the largest sum = %d\n", numbers, sum)

	numbers = []int{5, 4, -1, 7, 8}
	sum = maxSubArray(numbers)
	fmt.Printf("The array %v has the largest sum = %d\n", numbers, sum)

	numbers = []int{5, 4, 1, -7, 8}
	sum = maxSubArray(numbers)
	fmt.Printf("The array %v has the largest sum = %d\n", numbers, sum)
}

func maxSubArray(nums []int) int {
	maxSum, curSum := nums[0], 0

	for i := 0; i < len(nums); i++ {
		if curSum < 0 {
			curSum = 0
		}
		curSum += nums[i]
		maxSum = max(maxSum, curSum)
	}
	return maxSum
}

func max (num1, num2 int) int {
	if num1 >= num2 {
		return num1
	}
	return num2
}
