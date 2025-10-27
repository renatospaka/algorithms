package main

import (
	"fmt"
)

func main() {
	test(12)
	test(13)
}

func test(n int) {
	sqrt := minSquares(n)
	fmt.Printf("Is %d a perfect square (SQRT)? %v\n", n, sqrt)
	fmt.Println("=================================")
	fmt.Println()
}

func minSquares(n int) int {
	if n <= 0 {
		return 0
	}

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			square := j * j
			if square > i {
				break
			}
			if dp[i-square]+1 < dp[i] {
				dp[i] = dp[i-square] + 1
			}
		}
	}

	return dp[n]
}
