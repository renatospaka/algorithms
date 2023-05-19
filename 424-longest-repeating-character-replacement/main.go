package main

import "fmt"

func main() {
	var (
		text     string
		times    int
		replaces int
	)

	text = "ABAB"
	times = 2
	replaces = characterReplacementFaster(text, times)
	fmt.Printf("The text %s is now %d long when replaced %d times\n", text, replaces, times)

	text = "AABABBA"
	times = 1
	replaces = characterReplacementFaster(text, times)
	fmt.Printf("The text %s is now %d long when replaced %d times\n", text, replaces, times)
}

func characterReplacement(text string, k int) int {
	hash := make(map[string]int)
	result := 0

	left, right, masxFreq := 0, 0, 0
	for i := 0; i < len(text); i++ {
		right = i
		char := string(text[right])
		hash[char]++
		masxFreq = max(masxFreq, hash[char])
		fmt.Printf("right: %d => text[right]: %v\n", right, char)

		if (right - left + 1) - masxFreq > k {
			char = string(text[left])
			hash[char]--
			left++
		}
	}
	
	result = max(result, right - left + 1)
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}


func characterReplacementFaster(text string, k int) int {
	asc := [26]int{}
	maxIndex := 0
	left := 0

	for i := 0; i < len(text); i++ {
		asc[text[i]-'A']++
		if asc[text[i]-'A'] > asc[text[maxIndex]-'A'] {
			maxIndex = i
		}

		if i - left + 1 - asc[text[maxIndex]-'A'] > k {
			asc[text[left]-'A']--
			left++
		}
	}
	return len(text) - left
}
