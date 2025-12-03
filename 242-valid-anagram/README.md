# 242. Valid Anagram

**Difficulty:** Easy

## Problem

Given two strings `s` and `t`, return `true` if `t` is an anagram of `s`, and `false` otherwise.

An anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

## Approach

Count the frequency of each character in both strings and compare:

1. If the lengths of `s` and `t` are different, return `false`.
2. Use a hash map or array to count the occurrences of each character in `s` and `t`.
3. Compare the frequency counts for both strings. If they match, return `true`; otherwise, return `false`.

**Time complexity:** $O(n)$  
**Space complexity:** $O(1)$ (since the alphabet size is fixed)

## Example

**Input:** s = "anagram", t = "nagaram"  
**Output:** true  
**Explanation:** Both strings have the same letters with the same frequency.

**Input:** s = "rat", t = "car"  
**Output:** false  
**Explanation:** The strings do not have the same letters.

## Usage

See `main.go` for code to check if two strings are anagrams.

## Resources

- [LeetCode problem](https://leetcode.com/problems/valid-anagram/)
- [Valid Anagram (YouTube)](https://youtu.be/9UtInBqnCgA?si=OCo8n2Diw-XQrSHj)