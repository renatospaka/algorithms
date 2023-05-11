package main

import "fmt"

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	station := calcGasStation(gas, cost)
	fmt.Printf("Initial station: %d\n", station)
}

func calcGasStation(gas []int, cost []int) int {
	station, total := -1, 0
	sumGas, sumCost := returnSumArray(gas, cost)
	if sumGas < sumCost {
		return station
	}

	for i := 0; i < len(gas); i++ {
		total += gas[i] - cost[i]
		if total < 0 {
			total = 0
			station = i + 1
		}
	}
	return station
}

func returnSumArray(gas []int, cost []int) (int, int) {
	sumGas := -1
	sumCost := -1

	for i := 0; i < len(gas); i++ {
		sumGas += gas[i]
		sumCost += cost[i]
	}

	return sumGas, sumCost
}
