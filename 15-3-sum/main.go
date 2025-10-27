package main

import (
	"fmt"
	"sort"
)

func main() {
	test([]int{-1, 0, 1, 2, -1, -4})
	test([]int{0, 1, 1})
	test([]int{0, 0, 0})
	test([]int{1, 5, -3, 0, -1, 2, 2, -2})
	test([]int{-2, 0, 1, 1, 2})
	test([]int{-2, 0, 0, 2, 2})
}

func test(nums []int) {
	result := threeSum(nums)
	fmt.Printf("The array that sums zero from %v is %v\n", nums, result)
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	size := len(nums)
	// fmt.Printf("sorted: %v\n", nums)
	var result [][]int
	for i := 0; i < size-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, size-1
		// innerSize := -1
		for left < right {
			sum := nums[left] + nums[right] + nums[i]
			// fmt.Printf("i: %d => left: %d || right: %d || nums[i]: %d || nums[left]: %d || nums[right]: %d || sum: %d\n", i, left, right, nums[i], nums[left], nums[right], sum)
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}

				for left < right && nums[right] == nums[right+1] {
					right--
				}
				// fmt.Printf("i: %d => sum: %d ==> result: %v\n", i, sum, result)

			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}
