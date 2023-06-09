package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var (
		text string
		isIt bool
	)

	text = "A man, a plan, a canal: Panama"
	isIt = isPalindrome2(text)
	fmt.Printf("The phrase <%s> is palindrome: %t\n", text, isIt)

	text = "Renato Tane R"
	isIt = isPalindrome2(text)
	fmt.Printf("The phrase <%s> is palindrome: %t\n", text, isIt)

	text = ".,"
	isIt = isPalindrome2(text)
	fmt.Printf("The phrase <%s> is palindrome: %t\n", text, isIt)
}

func isPalindrome(text string) bool {
	left, right := 0, len(text) - 1
	for left < right {
		ascL :=  text[left]
		letterL :=  strings.ToLower(string(ascL))
		// fmt.Printf("letterL: %s, ascL: %d, left: %d\n", letterL, ascL, left)
		for !isAlphanumeric(letterL) && left < right {
			left++
			ascL =  text[left]
			letterL =  strings.ToLower(string(ascL))
			// fmt.Printf("letterL: %s, ascL: %d, left: %d\n", letterL, ascL, left)
		}

		ascR := text[right]
		letterR :=  strings.ToLower(string(ascR))
		// fmt.Printf("letterR: %s, ascR: %d, right: %d\n", letterR, ascR, right)
		for !isAlphanumeric(letterR) && left < right {
			right--
			ascR = text[right]
			letterR =  strings.ToLower(string(ascR))
			// fmt.Printf("letterR: %s, ascR: %d, right: %d\n", letterR, ascR, right)
		}

		if letterL != letterR {
			return false
		}
		left++
		right--
	} 
	return true
}

func isAlphanumeric(letter string) bool {
	asc := letter[0]
	return (asc >= 97 && asc <= 122) ||
	(asc >= 65 && asc <= 90) ||
	(asc >= 48 && asc <= 57)
}




func isPalindrome2(text string) bool {
	left, right := 0, len(text) - 1
	hash := []rune(text)

	for left < right {
		leftCh := unicode.ToLower(hash[left])
		if !isAlphanumeric2(leftCh) {
			left++
			continue
		}
		
		rightCh := unicode.ToLower(hash[right])
		if !isAlphanumeric2(rightCh) {
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

func isAlphanumeric2(letter rune) bool {
	return unicode.IsLetter(letter) || unicode.IsDigit(letter)
}