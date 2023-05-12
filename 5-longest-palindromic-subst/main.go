package main

import "fmt"

func main() {
	text := "babad"
	palidrom := checkLongestPalindromic(text)
	fmt.Printf("palidrom for text: %s is %s\n", text, palidrom)
}

func checkLongestPalindromic(text string) string {
	palidrom := ""
	size := 0

	for i := 0; i < len(text) - 1; i++ {
		left, right := i, i 
		leftS := string(text[left])
		rightS := string(text[right])
		fmt.Printf("i: %d, left: %d, leftS: %s, right: %d, rightS: %s\n", i, left, leftS, right, rightS)
		for left >= 0 && right < len(text) && leftS == rightS {
			if (right - left + 1) > size {
				palidrom = text[left:right + 1]
				size = right - left + 1
				fmt.Println("palidrom:", palidrom)
			}
			left--
			right++
		}

		left, right = i, i + 1
		leftS = string(text[left])
		rightS = string(text[right])
		fmt.Printf("i: %d, left: %d, leftS: %s, right: %d, rightS: %s\n", i, left, leftS, right, rightS)
		for left >= 0 && right < len(text) && leftS == rightS {
			if (right - left + 1) > size {
				palidrom = text[left:right + 1]
				size = right - left + 1
				fmt.Println("palidrom:", palidrom)
			}
			left--
			right++
		}
	}

	return palidrom
}
