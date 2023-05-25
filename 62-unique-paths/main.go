package main

import "fmt"

func main() {
	var (
		horizontal int
		vertical   int
		count      int
	)

	vertical, horizontal = 3, 7
	count = uniquePaths(vertical, horizontal)
	fmt.Printf("For a grid of %dx%d there are %d ways to achieve the lower right corner\n", vertical, horizontal, count)

	vertical, horizontal = 3, 2
	count = uniquePaths(vertical, horizontal)
	fmt.Printf("For a grid of %dx%d there are %d ways to achieve the lower right corner\n", vertical, horizontal, count)

	vertical, horizontal = 1, 1
	count = uniquePaths(vertical, horizontal)
	fmt.Printf("For a grid of %dx%d there are %d ways to achieve the lower right corner\n", vertical, horizontal, count)

	vertical, horizontal = 1, 2
	count = uniquePaths(vertical, horizontal)
	fmt.Printf("For a grid of %dx%d there are %d ways to achieve the lower right corner\n", vertical, horizontal, count)

	vertical, horizontal = 2, 2
	count = uniquePaths(vertical, horizontal)
	fmt.Printf("For a grid of %dx%d there are %d ways to achieve the lower right corner\n", vertical, horizontal, count)
}

func uniquePaths(vertical int, horizontal int) int {
	dp := make([]int, horizontal)
	dp[0] = 1

	for row := 0; row < vertical; row++ {
		for col := 1; col < horizontal; col++ {
			dp[col] += dp[col - 1]
		}
	}

	return dp[horizontal - 1]
}
