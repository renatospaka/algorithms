package main

import "fmt"

func main() {
	var (
		text1 string
		text2 string
		anagram bool
	)

	text1 = "anagram"
	text2 = "nagaram"
	anagram = isAnagram(text1, text2)
	fmt.Printf("%s and %s are anagrams? %t\n", text1, text2, anagram)

	text1 = "rat"
	text2 = "cat"
	anagram = isAnagram(text1, text2)
	fmt.Printf("%s and %s are anagrams? %t\n", text1, text2, anagram)

	text1 = "c"
	text2 = "c"
	anagram = isAnagram(text1, text2)
	fmt.Printf("%s and %s are anagrams? %t\n", text1, text2, anagram)
}

func isAnagram(text1 string, text2 string) bool {
	if len(text1) < 1 {
		return false
	}
	if len(text1) != len(text2) {
		return false
	}

	anag1 := make(map[string]int)
	anag2 := make(map[string]int)

	for i := 0; i < len(text1); i++ {
		letter1 := string(text1[i])
		letter2 := string(text2[i])
		// fmt.Printf("i: %d, letter1: %s\n", i, letter1)

		anag1[letter1]++
		anag2[letter2]++
		// fmt.Printf("i: %d -> letter1: %s, anag1[letter1]: %d, letter2: %s, anag2[letter2]: %d\n", i, letter1, anag1[letter1], letter2, anag2[letter2])
	}

	// fmt.Printf("anag2[letter2]: %d itens, anag1[letter1]: %d items\n", len(anag1), len(anag2))
	if len(anag1) != len(anag2) {
		return false
	}

	for letter := range anag1 {
		duplicates := anag1[letter] == anag2[letter]
		// fmt.Println("letter", letter, "value", anag1[letter], "value2", anag2[letter], "duplicates", duplicates)
		if !duplicates {
			return false
		}
	}
	return true
}
