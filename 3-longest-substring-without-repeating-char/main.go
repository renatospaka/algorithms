package main

import (
	"fmt"
	"strings"
)

func main() {
	longest := lengthOfLongestSubstring("abcabcbb")
	fmt.Printf("testing 'abcabcbb': %d\n", longest)
	
	longest = lengthOfLongestSubstring("bbbbb")
	fmt.Printf("testing 'bbbbb': %d\n", longest)

	longest = lengthOfLongestSubstring("pwwkew")
	fmt.Printf("testing 'pwwkew': %d\n", longest)

	longest = lengthOfLongestSubstring("")
	fmt.Printf("testing '': %d\n", longest)

}

func lengthOfLongestSubstring(text string) int {
	longest := 0
	longestTemp := ""
	for _, letter := range text {
		ix := strings.Index(longestTemp, string(letter))
		// fmt.Printf("Letter: %d (%s) => %d",letter, string(letter), ix)
		if ix != -1 {
			if len(longestTemp) > longest {
				longest = len(longestTemp)
			}
			longestTemp = longestTemp[ix + 1:]
			// fmt.Printf(" ==> %s", longestTemp)
		}
		longestTemp += string(letter)
	}

	if len(longestTemp) > longest {
		longest = len(longestTemp)
	}

	return longest
}
