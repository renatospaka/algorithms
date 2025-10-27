package main

import (
	"fmt"
)

func main() {
	test(16)
	test(14)
}

func test(n int) {
	sqrt := isPerfectSquareSQRT(n)
	logN := isPerfectSquareLogN(n)
	fmt.Printf("Is %d a perfect square (SQRT)? %v\n", n, sqrt)
	fmt.Printf("Is %d a perfect square (LogN)? %v\n", n, logN)
	fmt.Println("=================================")
	fmt.Println()
}

func isPerfectSquareSQRT(n int) bool {
	if n <= 0 {
		return false
	}

	for i := 1; i+1 <= n; i++ {
		if i*i == n {
			return true
		}
		if i*i > n {
			return false
		}
	}
	return false
}

func isPerfectSquareLogN(n int) bool {
	if n <= 0 {
		return false
	}

	left, right := 1, n
	for left <= right {
		mid := left + (right-left)/2
		squared := mid * mid

		if squared == n {
			return true
		} else if squared < n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
