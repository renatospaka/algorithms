package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	test("A man, a plan, a canal: Panama")
	test("race a car")
	test("Renato Tane R")
	test(".,")
}

func test(s string) {
	palindrome := isPalindrome(s)
	fmt.Printf("The phrase <%s> is palindrome: %t", s, palindrome)
	faster := isPalindromeFaster(s)
	fmt.Printf(" | is palindrome (faster): %t\n", faster)
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return false
	}
	if len(s) == 1 {
		return true
	}

	left, right := 0, len(s)-1
	for left < right {
		ascL := s[left]
		letterL := strings.ToLower(string(ascL))
		for !isAlphanumeric(letterL) && left < right {
			left++
			ascL = s[left]
			letterL = strings.ToLower(string(ascL))
		}

		ascR := s[right]
		letterR := strings.ToLower(string(ascR))
		for !isAlphanumeric(letterR) && left < right {
			right--
			ascR = s[right]
			letterR = strings.ToLower(string(ascR))
		}

		if letterL != letterR {
			return false
		}
		left++
		right--
	}
	return true
}

func isPalindromeFaster(s string) bool {
	if len(s) == 0 {
		return false
	}
	if len(s) == 1 {
		return true
	}

	left, right := 0, len(s)-1
	hash := []rune(s)

	for left < right {
		leftCh := unicode.ToLower(hash[left])
		if !isAlphanumericLibrary(leftCh) {
			left++
			continue
		}

		rightCh := unicode.ToLower(hash[right])
		if !isAlphanumericLibrary(rightCh) {
			right--
			continue
		}

		if leftCh != rightCh {
			return false
		}
		left++
		right--
	}

	return true
}

func isAlphanumeric(letter string) bool {
	asc := letter[0]
	// 48 - 57 ==> 0 to 9
	// 65 - 90 ==> A to Z
	// 97 - 122 ==> a to z
	return (asc >= 97 && asc <= 122) ||
		(asc >= 65 && asc <= 90) ||
		(asc >= 48 && asc <= 57)
}

func isAlphanumericLibrary(letter rune) bool {
	return unicode.IsLetter(letter) || unicode.IsDigit(letter)
}
