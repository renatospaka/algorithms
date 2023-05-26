package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		numbers [][]int
		removed int
	)

	numbers = [][]int{{1,2},{2,3},{3,4},{1,3}}
	removed = eraseOverlapIntervals(numbers)
	fmt.Printf("%d interval/s can be removed from %v\n", removed, numbers)

	numbers = [][]int{{1,2},{1,2},{1,2}}
	removed = eraseOverlapIntervals(numbers)
	fmt.Printf("%d interval/s can be removed from %v\n", removed, numbers)

	numbers = [][]int{{1,2},{2,3}}
	removed = eraseOverlapIntervals(numbers)
	fmt.Printf("%d interval/s can be removed from %v\n", removed, numbers)
}

// greedy?
// sort by start
func eraseOverlapIntervals(intervals [][]int) int {
	const (
		start	= 0 
		end 	= 1
	) 
	sort.Slice(intervals, func(i, j int) bool {return intervals[i][start] < intervals[j][start]})

	removed, currEnd := 0, intervals[0][end]
	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		if current[start] < currEnd {
			removed++
			currEnd = min(current[end], currEnd)

		} else {
			currEnd = current[end]
		}
	}

	return removed
}

func min (num1, num2 int) int {
	if num1 <= num2 {
		return num1
	}
	return num2
}
