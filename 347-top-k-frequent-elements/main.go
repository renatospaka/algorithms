package main

import "fmt"

func main() {
	var(
		list []int
		frequency []int
		factor int
	)

	factor = 2
	list = []int{1,1,1,2,2,3}
	frequency = topKFrequentSuccess(list, factor)
	fmt.Printf("The %dst frequent numbers of %v are %v\n", factor, list, frequency)
	
	factor = 1
	list = []int{1}
	frequency = topKFrequentSuccess(list, factor)
	fmt.Printf("The %dst frequent numbers of %v are %v\n", factor, list, frequency)

	factor = 4
	list = []int{1,1,1,2,2,3,3,3,3,4,4,5,6,7,8}
	frequency = topKFrequentSuccess(list, factor)
	fmt.Printf("The %dst frequent numbers of %v are %v\n", factor, list, frequency)
}

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	frequency := make([]int, len(nums) - 1)

	for i := 0; i < len(nums); i++ {
		count[nums[k]] += 1
	}

	for num, total := range count {
		frequency[num] = total
	}

	return frequency
}

func topKFrequentSuccess(nums []int, k int) []int {
	countMap := map[int]int{}
	countSlice := make([][]int, len(nums) + 1)

	for _, num := range nums {
		if count, ok := countMap[num]; ok {
			countMap[num] = count + 1
		} else {
			countMap[num] = 1
		}
	}

	for num, count := range countMap {
		countSlice[count] = append(countSlice[count], num)
	}

	result := []int{}
	for i := len(countSlice) - 1; i > 0; i-- {
		result = append(result, countSlice[i]...)
		if len(result) == k {
			return result
		}
	}
	return result
}
