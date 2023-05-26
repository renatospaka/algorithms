package main

import "fmt"

func main() {
	var (
		numbers         [][]int
		newNumbers      []int
		newNumbersOrig  []int
		insertedNumbers [][]int
	)

	numbers = [][]int{{1,3},{6,9}}
	newNumbers = []int{2,5}
	newNumbersOrig = []int{2,5}
	insertedNumbers = insert(numbers, newNumbers)
	fmt.Printf("The original interval %v has changed to %v after new interval %v was inserted\n", numbers, insertedNumbers, newNumbersOrig)

	numbers = [][]int{{1,2},{3,5},{6,7},{8,10},{12,16}}
	newNumbers = []int{4, 8}
	newNumbersOrig = []int{4, 8}
	insertedNumbers = insert(numbers, newNumbers)
	fmt.Printf("The original interval %v has changed to %v after new interval %v was inserted\n", numbers, insertedNumbers, newNumbersOrig)

	numbers = [][]int{{1,3},{6,9}}
	newNumbers = []int{2,7}
	newNumbersOrig = []int{2,7}
	insertedNumbers = insert(numbers, newNumbers)
	fmt.Printf("The original interval %v has changed to %v after new interval %v was inserted\n", numbers, insertedNumbers, newNumbersOrig)
}

func insert(intervals [][]int, newInterval []int) [][]int {
	response := [][]int{}

	for pos, currInt := range intervals {
		currStart, currEnd := currInt[0], currInt[1]
		newIntStart, newIntEnd := newInterval[0], newInterval[1]
		if newIntEnd < currStart {
			response = append(response, newInterval)
			response = append(response, intervals[pos:]...)
			return response

		} else
		if newIntStart > currEnd {
			response = append(response, currInt)

		} else {
			newInterval[0] = min(newIntStart, currStart)
			newInterval[1] = max(newIntEnd, currEnd)
		}
	}

	response = append(response, newInterval)
	return response
}

func max (num1, num2 int) int {
	if num1 >= num2 {
		return num1
	}
	return num2
}

func min (num1, num2 int) int {
	if num1 <= num2 {
		return num1
	}
	return num2
}
