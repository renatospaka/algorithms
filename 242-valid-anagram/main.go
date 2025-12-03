package main

import "fmt"

func main() {
	test("anagram", "nagaram")
	test("rat", "cat")
	test("c", "c")
	test("", "")
	test("listen", "silent")
	test("vile", "evil")
}

func test(s string, t string) {
	anagram := isAnagram(s, t)
	fmt.Printf("'%s' and '%s' are anagrams? %t", s, t, anagram)
	faster := isAnagramFaster(s, t)
	fmt.Printf(" | using (faster)? %t\n", faster)
}

func isAnagram(s string, t string) bool {
	if len(s) < 1 {
		return false
	}
	if len(s) != len(t) {
		return false
	}

	anag1 := make(map[string]int, len(s))
	anag2 := make(map[string]int, len(t))

	for i := 0; i < len(s); i++ {
		letter1 := string(s[i])
		anag1[letter1]++
		letter2 := string(t[i])
		anag2[letter2]++
	}

	for letter := range anag1 {
		if duplicates := anag1[letter] == anag2[letter]; !duplicates {
			return false
		}
	}
	return true
}

func isAnagramFaster(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := [26]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']--
		count[t[i]-'a']++
	}

	for _, c := range count {
		if c != 0 {
			return false
		}
	}
	return true
}
