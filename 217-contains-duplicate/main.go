package main

import "fmt"

func main() {

	test([]int{1, 2, 3, 1})
	test([]int{1, 2, 3, 4})
	test([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})
	test([]int{3, 3})
	test([]int{2})
	test([]int{})
}

func test(list []int) {
	duplicate := containDuplicates(list)
	fmt.Printf("The array %v contains duplicates? %t\n", list, duplicate)
}

func containDuplicates(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	hash := make(map[int]struct{}, len(nums))
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if _, exists := hash[num]; exists {
			return true
		}
		hash[num] = struct{}{}
	}

	return false
}
