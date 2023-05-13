package main

import "fmt"

func main() {

	var (
		prices []int
		bestDay int
	)

	prices = []int{7, 1, 5, 3, 6, 4}
	bestDay = maxProfit(prices)
	fmt.Printf("The maximum profit for %v is to sell in %d\n", prices, bestDay)

	prices = []int{7, 6, 4, 3, 1}
	bestDay = maxProfit(prices)
	fmt.Printf("The maximum profit for %v is to sell in %d\n", prices, bestDay)

	prices = []int{7, 7}
	bestDay = maxProfit(prices)
	fmt.Printf("The maximum profit for %v is to sell in %d\n", prices, bestDay)

	prices = []int{}
	bestDay = maxProfit(prices)
	fmt.Printf("The maximum profit for %v is to sell in %d\n", prices, bestDay)
}

func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	maxProfit := 0
	lowest := prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] < lowest {
			lowest = prices[i]
		}

		// fmt.Printf("i: %d, maxProfit: %d, prices[i]: %d, lowest: %d\n", i, maxProfit, prices[i], lowest)
		if maxProfit < prices[i] - lowest {
			maxProfit = prices[i] - lowest
		}
	}

	return maxProfit
}
