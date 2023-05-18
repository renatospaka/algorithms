package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var (
		strs []string
		anagrams [][]string
	)

	strs = []string{"eat","tea","tan","ate","nat","bat"}
	anagrams = groupAnagrams(strs)
	fmt.Printf("Anagrams for %v are (if any): %v\n", strs, anagrams)

	strs = []string{"a"}
	anagrams = groupAnagrams(strs)
	fmt.Printf("Anagrams for %v are (if any): %v\n", strs, anagrams)

	strs = []string{""}
	anagrams = groupAnagrams(strs)
	fmt.Printf("Anagrams for %v are (if any): %v\n", strs, anagrams)

	strs = []string{"rato","tora","rota","otra","atri","gato"}
	anagrams = groupAnagrams(strs)
	fmt.Printf("Anagrams for %v are (if any): %v\n", strs, anagrams)
}

func groupAnagrams(strs []string) [][]string {
  words := make(map[string][]string)
	for _, word := range strs {
		key := sorting(word)
		if _, ok := words[key]; !ok {
			words[key] = []string{}
		}
		words[key] = append(words[key], word)
		// fmt.Printf("word: %s | words[key]: %v | key: %s\n", word, words[key], key)
	}

	result := make([][]string, len(words))
	ix := 0
	for _, anagrams := range words {
		result[ix] = anagrams
		ix++
	}
	return result
}

func sorting(str string) string {
	strArr := strings.Split(str, "")
	sort.Strings(strArr)
	return strings.Join(strArr, "")
}
