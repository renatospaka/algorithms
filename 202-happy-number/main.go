package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		number int
		happy  bool
	)

	number = 19
	happy = isHappy(number)
	fmt.Printf("The number %d is happy? %t\n", number, happy)

	number = 2
	happy = isHappy(number)
	fmt.Printf("The number %d is happy? %t\n", number, happy)

	number = 13
	happy = isHappy(number)
	fmt.Printf("The number %d is happy? %t\n", number, happy)
}

func isHappy(number int) bool {
	existing := make(map[int]bool)
	total := 0

	for {
		if exists := existing[number]; exists {
			break
		}

		existing[number] = true
		toStr := fmt.Sprint(number)
		size := len(toStr)

		for i := 0; i < size; i++ {
			digit, _ := strconv.Atoi(string(toStr[i]))
			total += digit * digit
		}

		if total == 1 {
			return true
		}

		number = total
		total = 0
	}
	return false
}
