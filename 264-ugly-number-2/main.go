package main

import (
	"fmt"
)

func main() {
	test(10)
	test(1)
}

func test(n int) {
	ugly := nthUgly(n)
	fmt.Printf("n th ugly number for %d is %d\n", n, ugly)
}

func nthUgly(n int) int {
	if n <= 0 {
		return -1
	}

	nums, i2, i3, i5 := []int{1}, 0, 0, 0

	for i := 1; i < n; i++ {
		next2, next3, next5 := nums[i2]*2, nums[i3]*3, nums[i5]*5
		nextUgly := min(next2, min(next3, next5))
		nums = append(nums, nextUgly)

		if nextUgly == next2 {
			i2++
		}
		if nextUgly == next3 {
			i3++
		}
		if nextUgly == next5 {
			i5++
		}
	}

	return nums[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
