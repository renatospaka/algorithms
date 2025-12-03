package main

import "fmt"

func main() {
	test(2)
	test(3)
	test(5)
}

func test(n int) {
	ways := climbStairsFibbonacci(n)
	fmt.Printf("There are %d ways to climb a stairs of %d n\n", ways, n)
}

// dynamic problem - caching solutions, memorization, bottom up (Fibbonacci Sequence) <=
func climbStairsFibbonacci(n int) int {
	if n < 3 {
		return n
	}

	// starting at nth position, nth and n-1th position receive value 1
	// because it takes 1 step (and only one) to go from n-1th to nth
	one, two := 1, 1
	for i := 0; i < n-1; i++ {
		temp := one
		one = one + two
		two = temp
	}
	return one
}
