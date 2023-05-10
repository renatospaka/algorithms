package main

import (
	"fmt"
	"strings"
)

func main() {
	longest := longestSubstringWithoutRepeatingChar("abcabcbb")
	fmt.Printf("testing 'abcabcbb': %d\n", longest)
	
	longest = longestSubstringWithoutRepeatingChar("bbbbb")
	fmt.Printf("testing 'bbbbb': %d\n", longest)

	longest = longestSubstringWithoutRepeatingChar("pwwkew")
	fmt.Printf("testing 'pwwkew': %d\n", longest)

	longest = longestSubstringWithoutRepeatingChar("")
	fmt.Printf("testing '': %d\n", longest)

}

func longestSubstringWithoutRepeatingChar(text string) int {
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
