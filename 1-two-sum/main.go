package main

import "fmt"

func main() {
	test([]int{2, 7, 11, 15}, 9)
	test([]int{3, 2, 4}, 6)
	test([]int{3, 3}, 6)
	test([]int{5}, 5)
	test([]int{2, 7, 11, 15}, 26)
}

func test(nums []int, target int) {
	indexes := twoSum(nums, target)
	fmt.Printf("The indexes that satisfy sum=%d of array %v is %v\n", target, nums, indexes)
}

func twoSum(nums []int, target int) []int {
	if len(nums) < 1 {
		return []int{}
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return nums
		}
		return []int{}
	}
	hash := make(map[int]int, len(nums))

	for i, num := range nums {
		if ix, exists := hash[target-num]; exists {
			return []int{ix, i}
		}
		hash[num] = i
	}
	return []int{}
}
