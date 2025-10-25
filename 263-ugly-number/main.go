package main

import (
	"fmt"
)

func main() {
	test(6)
	test(1)
	test(14)
}

func test(n int) {
	ugly := isUgly(n)
	fmt.Printf("Is %d ugly? %t\n", n, ugly)
}

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	for _, prime := range []int{2, 3, 5} {
		for n%prime == 0 {
			n /= prime
		}
	}
	return n == 1
}
