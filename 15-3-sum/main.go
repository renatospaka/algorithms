package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		unsorted []int
		sorted   []int
		sumOf    [][]int
	)

	unsorted = []int{-1, 0, 1, 2, -1, -4}
	sorted = []int{-1, 0, 1, 2, -1, -4}
	sumOf = threeSum(sorted)
	fmt.Printf("The array that sums zero from %v is %v\n", unsorted, sumOf)

	unsorted = []int{0, 1, 1}
	sorted = []int{0, 1, 1}
	sumOf = threeSum(sorted)
	fmt.Printf("The array that sums zero from %v is %v\n", unsorted, sumOf)

	unsorted = []int{0, 0, 0}
	sorted = []int{0, 0, 0}
	sumOf = threeSum(sorted)
	fmt.Printf("The array that sums zero from %v is %v\n", unsorted, sumOf)

	unsorted = []int{1, 5, -3, 0, -1, 2, 2, -2}
	sorted = []int{1, 5, -3, 0, -1, 2, 2, -2}
	sumOf = threeSum(sorted)
	fmt.Printf("The array that sums zero from %v is %v\n", unsorted, sumOf)

	unsorted = []int{-2, 0, 1, 1, 2}
	sorted = []int{-2, 0, 1, 1, 2}
	sumOf = threeSum(sorted)
	fmt.Printf("The array that sums zero from %v is %v\n", unsorted, sumOf)

	unsorted = []int{-2, 0, 0, 2, 2}
	sorted = []int{-2, 0, 0, 2, 2}
	sumOf = threeSum(sorted)
	fmt.Printf("The array that sums zero from %v is %v\n", unsorted, sumOf)
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	size := len(nums)
	// fmt.Printf("sorted: %v\n", nums)
	var result [][]int
	for idx := 0; idx < size - 2; idx++ {
		if idx > 0 && nums[idx] == nums[idx - 1] {
			continue
		} 

		left, right := idx + 1, size - 1
		// innerSize := -1
		for left < right {	
			sum := nums[left] + nums[right] + nums[idx]
			// fmt.Printf("idx: %d => left: %d || right: %d || nums[idx]: %d || nums[left]: %d || nums[right]: %d || sum: %d\n", idx, left, right, nums[idx], nums[left], nums[right], sum)
			if sum == 0 {
				result = append(result, []int{nums[idx], nums[left], nums[right]})
				left++
				right--

				for left < right && nums[left] == nums[left - 1] {
					left++
				}

				for left < right && nums[right] == nums[right + 1] {
					right--
				}
				// fmt.Printf("idx: %d => sum: %d ==> result: %v\n", idx, sum, result)

			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}
