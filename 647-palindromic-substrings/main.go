package main

import "fmt"

func main() {
	test("abc")
	test("aaa")
	test("aaab")
	test("b")
	test("abccbaabcda")
}

func test(s string) {
	count := countSubstrings(s)
	fmt.Printf("The s %s contains %d palindromic strings on it\n", s, count)
}

func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}

	qtd := 0
	for i := 0; i < len(s); i++ {
		qtd += count(s, i, i)   // testing for even palindromes
		qtd += count(s, i, i+1) // testing for odd palindromes
	}

	return qtd
}

func count(s string, left, right int) int {
	qtd := 0
	for left >= 0 && right < len(s) && s[left] == s[right] {
		qtd++
		left--
		right++
	}
	return qtd
}
