package main

import "fmt"

func main() {
	var (
		bits uint32
		ones int
	)

	bits = 00000000000000000000000000001011
	ones = hammingWeightBestPerformance(bits)
	fmt.Printf("The number of %b has %d '1'\n", bits, ones)

	bits = 00000000000000000000000010000000
	ones = hammingWeightBestPerformance(bits)
	fmt.Printf("The number of %b has %d '1'\n", bits, ones)

	// bits = 11111111111111111111111111111101
	// ones = hammingWeightBestPerformance(bits)
	// fmt.Printf("The number of %b has %d '1'\n", bits, ones)
}

func hammingWeight(number uint32) int {
	count := 0
	for number != 0 {
		number &= number - 1
		count += 1
	}
	return int(count)
}

func hammingWeightBestPerformance(number uint32) int {
	count := 0
	for ; number != 0; number /= 2 {
		if number % 2 == 1 {
			count++
		}
	}
	return count
}
