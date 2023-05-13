package main

import (
	"fmt"
	"math"
)

func main() {

	var (
		prices []int
		bestDay int
	)

	prices = []int{7, 1, 5, 3, 6, 4}
	bestDay = maxProfit3(prices)
	fmt.Printf("The maximum profit for %v is to sell in %d\n", prices, bestDay)

	prices = []int{7, 6, 4, 3, 1}
	bestDay = maxProfit3(prices)
	fmt.Printf("The maximum profit for %v is to sell in %d\n", prices, bestDay)

	prices = []int{7, 7}
	bestDay = maxProfit3(prices)
	fmt.Printf("The maximum profit for %v is to sell in %d\n", prices, bestDay)

	prices = []int{}
	bestDay = maxProfit3(prices)
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

func maxProfit2(prices []int) int {
	min := math.MaxUint32
	maxProfit := 0

	for _, price := range prices {
		if price > min {
			if price - min > maxProfit {
				maxProfit = price - min
			}
		} else {
			min = price
		}
	}

	return maxProfit
}

func maxProfit3(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	maxProfit := 0
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[1] < minPrice {
			minPrice = prices[i]
		}

		if prices[i] - minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit 
}
