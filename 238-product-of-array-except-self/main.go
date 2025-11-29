package main

import "fmt"

func main() {
	test([]int{1, 2, 3, 4})
	test([]int{-1, 1, 0, -3, 3})
}

func test(nums []int) {
	result := productExceptSelf(nums)
	fmt.Printf("The product of array except self for %v is %v\n", nums, result)
}

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	prefix := 1
	for i := 0; i < len(nums); i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}
