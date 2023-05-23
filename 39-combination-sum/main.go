package main

import "fmt"

func main() {
	var (
		candidates   []int
		target       int
		combinations [][]int
	)

	candidates = []int{2, 3, 6, 7}
	target = 7
	combinations = combinationSum(candidates, target)
	fmt.Printf("Possible combinatios of %v that sum %d are (if any): %v\n", candidates, target, combinations)

	candidates = []int{2, 3, 5}
	target = 8
	combinations = combinationSum(candidates, target)
	fmt.Printf("Possible combinatios of %v that sum %d are (if any): %v\n", candidates, target, combinations)

	candidates = []int{2}
	target = 1
	combinations = combinationSum(candidates, target)
	fmt.Printf("Possible combinatios of %v that sum %d are (if any): %v\n", candidates, target, combinations)
}

func combinationSum(candidates []int, target int) [][]int {
	current := make([]int, 0)
	combinations := make([][]int, 0)

	var dfs func(position int, sum int, current []int) // <- needs to be like this due to recursion
	dfs = func(position int, sum int, current []int) {
		if sum == target {
			// fmt.Printf(" %d = %d\n", sum, target)
			combinations = append(combinations, append([]int{}, current...)) // <- this is the catch
			return
		}

		if sum > target {
			// fmt.Printf(" %d >> %d\n", sum, target)
			return
		}

		for i := position; i < len(candidates); i++ {
			current = append(current, candidates[i])
			// fmt.Printf("position: %d || i: %d || candidates[i]: %d => %v\n", position, i, candidates[i], current)
			dfs(i, sum+candidates[i], current)
			current = current[:len(current) - 1] // remove/pop the last item from slice
			// fmt.Printf("**position: %d || i: %d || candidates[i]: %d => %v\n", position, i, candidates[i], current)
		}
	}

	dfs(0, 0, current)
	return combinations
}
