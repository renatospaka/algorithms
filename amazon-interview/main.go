package main

import "fmt"

func printComplementPairsOn2(angles []int) {
	for i := 0; i < len(angles); i++ {
		for j := i + 1; j < len(angles); j++ {
			if angles[i]+angles[j] == 90 {
				fmt.Printf("O(n*n) Pair: (%d, %d) => (%d, %d)\n", i, j, angles[i], angles[j])
			}
		}
	}
}

func printComplementPairsOn(angles []int) {
	seen := make(map[int][]int)
	for i, angle := range angles {
		complement := 90 - angle
		if indices, found := seen[complement]; found {
			for _, j := range indices {
				fmt.Printf("O(n)   Pair: (%d, %d) => (%d, %d)\n", j, i, angles[j], angle)
			}
		}
		seen[angle] = append(seen[angle], i)
	}
}

func main() {
	angles := []int{10, 40, 89, 23, 1, 50, 67, 80, 40, 50}
	printComplementPairsOn(angles)
	printComplementPairsOn2(angles)
}
