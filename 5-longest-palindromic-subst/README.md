# 5. Longest Palindromic Substring

**Difficulty:** Medium

## Problem

Given a string `s`, return the longest palindromic substring in `s`.

## Approach

Expand Around Center:

- Iterate through each character in the string, treating each as the center of a possible palindrome.
- For each center, expand outwards to check for both odd and even length palindromes.
- Track the longest palindromic substring found during the process.

Time complexity: $O(n^2)$  
Space complexity: $O(1)$

## Example

Input: s = "babad"  
Output: "bab"  
Explanation: "aba" is also a valid answer.

Input: s = "cbbd"  
Output: "bb"

Input: s = "anabanana"  
Output: "anana"

## Usage

See `main.go` for code to compute the longest palindromic substring in a given string.

## Resources

- [LeetCode problem](https://leetcode.com/problems/longest-palindromic-substring/)
- [Longest Palindromic Substring explanation (YouTube)](https://youtu.be/XYQecbcd6_c?si=Khynah-8PJ16m-08)
