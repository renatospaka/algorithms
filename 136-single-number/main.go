package main

import "fmt"

func main() {
	var (
		list   []int
		single int
	)

	list = []int{2, 2, 1}
	single = singleNumberMoreEfficient(list)
	fmt.Printf("The single number of %v is %d\n", list, single)

	list = []int{4, 1, 2, 1, 2}
	single = singleNumberMoreEfficient(list)
	fmt.Printf("The single number of %v is %d\n", list, single)

	list = []int{2, 1, 4, 1, 2}
	single = singleNumberMoreEfficient(list)
	fmt.Printf("The single number of %v is %d\n", list, single)

	list = []int{3, 5, 5, 3, 8}
	single = singleNumberMoreEfficient(list)
	fmt.Printf("The single number of %v is %d\n", list, single)

	list = []int{1}
	single = singleNumberMoreEfficient(list)
	fmt.Printf("The single number of %v is %d\n", list, single)
}

func singleNumber(list []int) int {
	result := 0
	for i := 0; i < len(list); i++ {
		result = list[i] ^ result
	}
	return result
}

func singleNumberMoreEfficient(list []int) int {
	for i := 1; i < len(list); i++ {
		list[0] ^= list[i]
	}
	return list[0]
}
