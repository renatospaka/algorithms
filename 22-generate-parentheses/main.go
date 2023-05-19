package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		size int
		result []string
	)

	size = 3
	result = generateParenthesis(size)
	fmt.Printf("Resoult of size of %d patenthesis: %v\n", size, result)

	size = 1
	result = generateParenthesis(size)
	fmt.Printf("Resoult of size of %d patenthesis: %v\n", size, result)

	size = 4
	result = generateParenthesis(size)
	fmt.Printf("Resoult of size of %d patenthesis: %v\n", size, result)
}

func generateParenthesis(n int) []string {
	var (
		stack []string
		result []string
		backTrack func(int, int)
	) 

	backTrack = func(openN, closeN int) {
		if openN == n && closeN == n && openN == closeN {
			result = append(result, strings.Join(stack, ""))
			return
		}

		if openN < n {
			stack = append(stack, "(")
			backTrack(openN + 1, closeN)
			removeLast(&stack)
		}

		if closeN < openN {
			stack = append(stack, ")")
			backTrack(openN, closeN + 1)
			removeLast(&stack)
		}
	}
	backTrack(0, 0)
	return result
}

func removeLast(list *[]string) {
	size := len(*list)
	*list = (*list)[:size - 1]
}
