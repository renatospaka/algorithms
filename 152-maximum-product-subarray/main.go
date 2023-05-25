package main

import "fmt"

func main() {
	var (
		numbers []int
		sum     int
	)

	numbers = []int{2, 3, -2, 4}
	sum = maxProduct(numbers)
	fmt.Printf("The maximum product subarray of %v is %d\n", numbers, sum)

	numbers = []int{-2, 0, -1}
	sum = maxProduct(numbers)
	fmt.Printf("The maximum product subarray of %v is %d\n", numbers, sum)

	numbers = []int{-3, -2, -1}
	sum = maxProduct(numbers)
	fmt.Printf("The maximum product subarray of %v is %d\n", numbers, sum)

	numbers = []int{-3, 2, -1}
	sum = maxProduct(numbers)
	fmt.Printf("The maximum product subarray of %v is %d\n", numbers, sum)
}

func maxProduct(nums []int) int {
	max, curMax, curMin := 0, 1, 1

	for i := 0; i < len(nums); i++ {
		temp := nums[i] * curMax
		curMax = maxOfThree(nums[i] * curMax, nums[i] * curMin, nums[i])
		curMin = minOfThree(temp, nums[i] * curMin, nums[i])
		max = maxOfTwo(max, curMax)

		// fmt.Printf("i: %d || nums[i]: %d || curMax: %d || curMin: %d || max: %d\n", i, nums[i], curMax, curMin, max)
	}
	return max
}

func maxOfTwo(num1, num2 int) int {
	if num1 >= num2{
		return num1
	}
	return num2
}

func maxOfThree(num1, num2, num3 int) int {
	if num1 >= num2 && num1 >= num3 {
		// fmt.Printf("maxOfThree(%d, %d, %d) is %d\n", num1, num2, num3, num1)
		return num1
		} else
	if num2 >= num1 && num2 >= num3 {
		// fmt.Printf("maxOfThree(%d, %d, %d) is %d\n", num1, num2, num3, num2)
		return num2
	}
	// fmt.Printf("maxOfThree(%d, %d, %d) is %d\n", num1, num2, num3, num3)
	return num3
}

func minOfThree(num1, num2, num3 int) int {
	if num1 <= num2 && num1 <= num3 {
		// fmt.Printf("minOfThree(%d, %d, %d) is %d\n", num1, num2, num3, num1)
		return num1
	} else
	if num2 <= num1 && num2 <= num3 {
		// fmt.Printf("minOfThree(%d, %d, %d) is %d\n", num1, num2, num3, num2)
		return num2
	}
	// fmt.Printf("minOfThree(%d, %d, %d) is %d\n", num1, num2, num3, num3)
	return num3
}
