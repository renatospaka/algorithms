package main

import "fmt"

func main() {
	test([]int{100, 4, 200, 1, 3, 2})
	test([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
	test([]int{1, 0, 1, 2})
	test([]int{0, 0})
}

func test(numbs []int) {
	result := longestConsecutive(numbs)
	fmt.Printf("The longest consecutive sequence in array %v is %v\n", numbs, result)
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return 1
	}

	numSet := make(map[int]struct{})
	for _, num := range nums {
		numSet[num] = struct{}{}
	}

	longest := 0
	for num := range numSet {
		if _, found := numSet[num-1]; !found {
			currentNum := num
			currentStreak := 1

			for {
				if _, found := numSet[currentNum+1]; found {
					currentNum++
					currentStreak++
				} else {
					break
				}
			}

			if currentStreak > longest {
				longest = currentStreak
			}
		}
	}
	return longest
}
