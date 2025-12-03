package main

import (
	"fmt"
	"math"
)

func main() {
	test([]int{7, 1, 5, 3, 6, 4})
	test([]int{7, 6, 4, 3, 1})
	test([]int{7, 7})
	test([]int{})
}

func test(prices []int) {
	bestDay := maxProfit(prices)
	fmt.Printf("The maximum profit for %v is to sell in %d\n", prices, bestDay)
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	maxProfit, minPrice := 0, prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[1] < minPrice {
			minPrice = prices[i]
		}

		if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}

func maxProfit3(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	maxProfit := 0
	lowest := prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] < lowest {
			lowest = prices[i]
		}

		// fmt.Printf("i: %d, maxProfit: %d, prices[i]: %d, lowest: %d\n", i, maxProfit, prices[i], lowest)
		if maxProfit < prices[i]-lowest {
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
			if price-min > maxProfit {
				maxProfit = price - min
			}
		} else {
			min = price
		}
	}

	return maxProfit
}
