package main

import "fmt"

func main() {
	var(
		number int
		result []int
	)

	number = 2
	result = countBits(number)
	fmt.Printf("The array for number = %d is %v\n", number, result)
}

// I don't know what, why and how...
func countBits(n int) []int {
	//dynamic programing 
	dp := make([]int, n + 1)
	offset := 1

	for i := 1; i <= n; i++ {
		if offset * 2 == i {
			offset = i
		}

		fmt.Printf("i: %d, offset: %d, dp[i]: %d\n", i, offset, dp[i])
		dp[i] = 1 + dp[i - offset]
	}
	return dp
}
