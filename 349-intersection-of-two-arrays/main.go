package main

import "fmt"

func main() {
	test([]int{1, 2, 2, 1}, []int{2, 2})
	test([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})
	test([]int{7, 1, 9, 7, 2}, []int{7, 2, 0})
}

func test(nums1 []int, nums2 []int) {
	result := intersection(nums1, nums2)
	fmt.Printf("The intersection of arrays %v and %v is %v\n", nums1, nums2, result)
}

func intersection(nums1 []int, nums2 []int) []int {
	seen := make(map[int]struct{})
	for _, n := range nums1 {
		seen[n] = struct{}{}
	}
	result := []int{}
	for _, n := range nums2 {
		if _, found := seen[n]; found {
			result = append(result, n)
			delete(seen, n) // Ensure uniqueness
		}
	}
	return result
}
