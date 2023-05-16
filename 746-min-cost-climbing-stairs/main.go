package main

import "fmt"

func main() {
	var(
		costs []int
		cost int
	)

	costs = []int{10, 15, 20}
	cost = minCostClimbingStairs(costs)
	fmt.Printf("The minimum costs to go upstairs with costs = %v is %d\n", costs, cost)

	costs = []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	cost = minCostClimbingStairs(costs)
	fmt.Printf("The minimum costs to go upstairs with costs = %v is %d\n", costs, cost)

	costs = []int{1, 100, 1, 1, 100, 100, 1, 1, 1, 1}
	cost = minCostClimbingStairs(costs)
	fmt.Printf("The minimum costs to go upstairs with costs = %v is %d\n", costs, cost)
}

func minCostClimbingStairs(costs []int) int {
	size := len(costs)
	if size == 2 {
		return min(costs[0], costs[1])
	}
	
	dp := make([]int, size + 1)
	for i := 2; i <= size; i++ {
		one := dp[i - 1] + costs[i - 1]
		two := dp[i - 2] + costs[i - 2]
		dp[i] = min(one, two)
	}
	return dp[size]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
