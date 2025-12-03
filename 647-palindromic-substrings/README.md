# 647. Palindromic Substrings

**Difficulty:** Medium

## Problem

Given a string `s`, return the number of palindromic substrings in it.

A substring is a contiguous sequence of characters within the string. A palindrome is a string that reads the same backward as forward.

Each substring must be counted only once, even if it appears multiple times in the string.

## Approach

Use the expand-around-center technique:

- Iterate through each character in the string.
- For each character, expand around it to find all palindromic substrings (odd length).
- Also expand between each pair of characters to find even-length palindromes.
- Count each valid palindrome found during expansion.

Time complexity: $O(n^2)$  
Space complexity: $O(1)$

## Example

Input: s = "abc"  
Output: 3  
Explanation: Three palindromic substrings: "a", "b", "c"

Input: s = "aaa"  
Output: 6  
Explanation: Six palindromic substrings: "a", "a", "a", "aa", "aa", "aaa"

## Usage

See `main.go` for code to compute the number of palindromic substrings in a given string.

## Resources

- [LeetCode problem](https://leetcode.com/problems/palindromic-substrings/)
- [Palindromic Substrings (YouTube)](https://youtu.be/4RACzI5-du8?si=5mHgFzv-PRYv79Rq)
