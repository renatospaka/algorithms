package main

import "fmt"

func main() {
	var (
		list []int
		target int
		indexes []int
	)
	
	target = 9
	list = []int{2, 7, 11 ,15}
	indexes = twoSum(list, target)
	fmt.Printf("The indexes that satisfy sum=%d of array %v is %v\n", target, list, indexes)
	
	target = 6
	list = []int{3, 2, 4}
	indexes = twoSum(list, target)
	fmt.Printf("The indexes that satisfy sum=%d of array %v is %v\n", target, list, indexes)
	
	target = 6
	list = []int{3, 3}
	indexes = twoSum(list, target)
	fmt.Printf("The indexes that satisfy sum=%d of array %v is %v\n", target, list, indexes)
	
	target = 5
	list = []int{5}
	indexes = twoSum(list, target)
	fmt.Printf("The indexes that satisfy sum=%d of array %v is %v\n", target, list, indexes)
}

func twoSum(list []int, target int) []int {
	if len(list) < 2 {
		return []int{}
	}
	hash := make(map[int]int)

	for ix, number := range list {
		if existingNumber, ok := hash[target - number]; ok {
			return []int{existingNumber, ix}
		}

		hash[number] = ix
	}

	return []int{}
}
