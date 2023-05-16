package main

import "fmt"

func main() {
	var(
		steps int
		ways int
	)

	steps = 2
	ways = climbStairs(steps)
	fmt.Printf("There are %d to climb a stars of %d steps\n", ways, steps)

	steps = 3
	ways = climbStairs(steps)
	fmt.Printf("There are %d to climb a stars of %d steps\n", ways, steps)

	steps = 5
	ways = climbStairs(steps)
	fmt.Printf("There are %d to climb a stars of %d steps\n", ways, steps)
}

// dynamic problem - caching solutions, memorization, bottom up (Fibbonacci Sequence) <=
func climbStairs(steps int) int {
	if steps < 3 {
		return steps
	}

	// starting at nth position, nth and n-1th position receive value 1 
	// because it takes 1 step (and only one) to go from n-1th to nth
	one, two := 1, 1
	for i := 0; i < steps - 1; i++ {
		temp := one
		one = one + two
		two = temp
	}
	return one
}
