package main

import "fmt"

func main() {
	test("babad")
	test("cbbd")
	test("a")
	test("anabanana")
	test("forgeeksskeegfor")
	test("abacdfgdcaba")
	test("abaxyzzyxf")
}

func test(s string) {
	longest := checkLongestPalindromic(s)
	fmt.Printf("longest longest for '%s' is %s\n", s, longest)
}

func checkLongestPalindromic(s string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s
	}

	longest := ""
	left, right, size := 0, 0, 0

	for i := 0; i < len(s)-1; i++ {
		left, right = i, i
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if (right - left + 1) > size {
				longest = s[left : right+1]
				size = right - left + 1
			}
			left--
			right++
		}

		left, right = i, i+1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if (right - left + 1) > size {
				longest = s[left : right+1]
				size = right - left + 1
			}
			left--
			right++
		}
	}

	return longest
}
