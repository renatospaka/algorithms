# 125. Valid Palindrome

**Difficulty:** Easy

## Problem

Given a string s, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

A palindrome is a word, phrase, or sequence that reads the same backward as forward, after removing non-alphanumeric characters and ignoring case.

## Approach

Use two pointers to check characters from both ends:

- Initialize two pointers, one at the start and one at the end of the string.
- Move the pointers towards each other, skipping non-alphanumeric characters.
- Compare the characters at each pointer (case-insensitive).
	- If they differ, return false.
	- If they match, continue.
- If all characters match, return true.

Time complexity: $O(n)$  
Space complexity: $O(1)$

## Example

Input: s = "A man, a plan, a canal: Panama"  
Output: true  
Explanation: "amanaplanacanalpanama" is a palindrome.

Input: s = "race a car"  
Output: false

Input: s = ".,"  
Output: true

## Usage

See `main.go` for code to check if a string is a valid palindrome.


## Resources

- [LeetCode problem](https://leetcode.com/problems/valid-palindrome/)
- [Valid Palindrome explanation (YouTube)](https://youtu.be/jJXJ16kPFWg?si=Eq411g1U85axM94a)
