package main

import "fmt"

func main() {
	var (
		numbers  []int
		missingN int
	)

	numbers = []int{3,0,1}
	missingN = missingNumber(numbers)
	fmt.Printf("The missing number in %v is %d\n", numbers, missingN)

	numbers = []int{9,6,4,2,3,5,7,0,1}
	missingN = missingNumber(numbers)
	fmt.Printf("The missing number in %v is %d\n", numbers, missingN)

	numbers = []int{0,1}
	missingN = missingNumber(numbers)
	fmt.Printf("The missing number in %v is %d\n", numbers, missingN)

	numbers = []int{0}
	missingN = missingNumber(numbers)
	fmt.Printf("The missing number in %v is %d\n", numbers, missingN)
}

func missingNumber(nums []int) int {
	missing := len(nums);

	for i := range nums {
		missing += i - nums[i]
	}
	return missing
}
