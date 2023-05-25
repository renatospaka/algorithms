package main

import "fmt"

func main() {
	var (
		word1   string
		word2   string
		longest int
	)

	word1, word2 = "abcde", "ace"
	longest = longestCommonSubsequence(word1, word2)
	fmt.Printf("The longest common subsequence of (%s/%s) has length of %d\n", word1, word2, longest)

	word1, word2 = "abc", "abc"
	longest = longestCommonSubsequence(word1, word2)
	fmt.Printf("The longest common subsequence of (%s/%s) has length of %d\n", word1, word2, longest)

	word1, word2 = "abc", "def"
	longest = longestCommonSubsequence(word1, word2)
	fmt.Printf("The longest common subsequence of (%s/%s) has length of %d\n", word1, word2, longest)

	word1, word2 = "chupa", "cabras"
	longest = longestCommonSubsequence(word1, word2)
	fmt.Printf("The longest common subsequence of (%s/%s) has length of %d\n", word1, word2, longest)
}

// create a 2D slice with the words and calculate bottom-up weather letters from one exist on the other word
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1) + 1)
	for i := 0; i < len(dp); i++ {
			dp[i] = make([]int, len(text2) + 1)
	}

	for i := len(text1) - 1; i >= 0; i-- {
		for j := len(text2) - 1; j >= 0; j-- {
			// walks diagonaly
			if text1[i] == text2[j] {
				dp[i][j] = 1 + dp[i + 1][j + 1]
			} else {
				dp[i][j] = max(dp[i][j + 1], dp[i + 1][j])
			}
		}
	}
	return dp[0][0]
}

func max(num1, num2 int) int {
	if num1 >= num2 {
		return num1
	}
	return num2
}