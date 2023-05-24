package main

import "fmt"

func main() {
	var (
		text 			string
		palidrom 	string
	)

	text = "babad"
	palidrom = checkLongestPalindromic(text)
	fmt.Printf("palidrom for text: %s is %s\n", text, palidrom)

	text = "cbbd"
	palidrom = checkLongestPalindromic(text)
	fmt.Printf("palidrom for text: %s is %s\n", text, palidrom)

	text = "anabanana"
	palidrom = checkLongestPalindromic(text)
	fmt.Printf("palidrom for text: %s is %s\n", text, palidrom)
}

func checkLongestPalindromic(text string) string {
	palidrom := ""
	left, right, size := 0, 0, 0

	for i := 0; i < len(text) - 1; i++ {
		left, right = i, i 
		fmt.Printf("i: %d, left: %d, text[left]: %v, right: %d, text[right]: %v\n", i, left, text[left], right, text[right])
		for left >= 0 && right < len(text) && text[left] == text[right] {
			if (right - left + 1) > size {
				palidrom = text[left:right + 1]
				size = right - left + 1
				// fmt.Println("palidrom:", palidrom)
			}
			left--
			right++
		}

		left, right = i, i + 1
		fmt.Printf("i: %d, left: %d, text[left]: %v, right: %d, text[right]: %v\n", i, left, text[left], right, text[right])
		for left >= 0 && right < len(text) && text[left] == text[right] {
			if (right - left + 1) > size {
				palidrom = text[left:right + 1]
				size = right - left + 1
				// fmt.Println("palidrom:", palidrom)
			}
			left--
			right++
		}
	}

	return palidrom
}
