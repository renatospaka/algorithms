package main

import "fmt"

func main() {
	var (
		text  string
		count int
	)

	text = "abc"
	count = countSubstrings(text)
	fmt.Printf("The text %s contains %d palindromic strings on it\n", text, count)

	text = "aaa"
	count = countSubstrings(text)
	fmt.Printf("The text %s contains %d palindromic strings on it\n", text, count)

	text = "b"
	count = countSubstrings(text)
	fmt.Printf("The text %s contains %d palindromic strings on it\n", text, count)

	text = "abccbaabcda"
	count = countSubstrings(text)
	fmt.Printf("The text %s contains %d palindromic strings on it\n", text, count)
}

func countSubstrings(text string) int {
	resp := 0
	for i := 0; i < len(text); i++ {
		resp += count(text, i, i)           // testing for even palindromes
		resp += count(text, i, i + 1)       // testing for odd palindromes
	}

	return resp
}

func count(text string, left int, right int) int {
	resp := 0
	for left >= 0 && right < len(text) && text[left] == text[right] {
		resp++
		left--
		right++
	}
	return resp
}