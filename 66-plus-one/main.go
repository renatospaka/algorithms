package main

import (
	"fmt"
)

func main() {
	var(
		numbers []int
		plus []int
	)

	numbers = []int{1, 2, 3}
	plus = plusOne(numbers)
	fmt.Printf("the array that representes %v is %v\n", []int{1, 2, 3}, plus)

	numbers = []int{1, 9, 9, 9}
	plus = plusOne(numbers)
	fmt.Printf("the array that representes %v is %v\n", []int{1, 9, 9, 9}, plus)

	numbers = []int{5, 4, 3, 2, 1}
	plus = plusOne(numbers)
	fmt.Printf("the array that representes %v is %v\n", []int{5, 4, 3, 2, 1}, plus)

	numbers = []int{9, 8}
	plus = plusOne(numbers)
	fmt.Printf("the array that representes %v is %v\n", []int{9, 8}, plus)
}

// func plusOneDraft(list []int) []int {
// 	newList := []int{}
// 	integer := 0

// 	for i := 0; i < len(list); i++ {
// 		// newList := newList
// 		pow := len(list) - i - 1
// 		tenth := int(math.Pow10(pow)) * list[i]

// 		integer += tenth
// 		newList = append(newList, list[i])
// 		if i == len(list) - 1 {
// 			integer++
// 			newList[i] = list[i] + 1
// 			if list[i] == 9 {
// 				// rebuild the array
// 			}
// 		}
// 		fmt.Println(i, tenth, integer)
// 	}
// 	fmt.Println(newList)
// 	return newList
// }

func plusOne(list []int) []int {
	for i := len(list) - 1; i >= 0 ; i-- {
		if list[i] < 9 {
			list[i] += 1
			return list
		}
		list[i] = 0
		// fmt.Println(" -> ", i, list[i])
	}
	list[0] = 1
	list = append(list, 0)
	return list
}
