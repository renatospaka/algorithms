# 217. Contains Duplicate

**Difficulty:** Easy

## Problem

Given an integer array `nums`, return `true` if any value appears at least twice in the array, and return `false` if every element is distinct.

## Approach

Use a hash set to track seen numbers:

1. Initialize an empty set.
2. Iterate through each number in `nums`:
	- If the number is already in the set, return `true`.
	- Otherwise, add the number to the set.
3. If no duplicates are found, return `false`.

**Time complexity:** $O(n)$  
**Space complexity:** $O(n)$

## Example

**Input:** nums = [1,2,3,1]  
**Output:** true  
**Explanation:** 1 appears twice.

**Input:** nums = [1,2,3,4]  
**Output:** false  
**Explanation:** All elements are distinct.

## Usage

See `main.go` for code to check if an array contains duplicates.

## Resources

- [LeetCode problem](https://neetcode.io/solutions/problems/contains-duplicate)
- [Contains Duplicate (YouTube)](https://youtu.be/3OamzN90kPg?si=OX5oHE0K7UOY9jc5)
