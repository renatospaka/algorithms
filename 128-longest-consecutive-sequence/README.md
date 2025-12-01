# 128. Longest Consecutive Sequence

**Difficulty:** Medium

## Problem

Given an unsorted array of integers `nums`, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in $O(n)$ time.

## Approach

Use a HashSet to achieve $O(n)$ time complexity:

1. Add all numbers to a set for fast lookup.
2. Iterate through each number in the array:
   - If the number is the start of a sequence (i.e., `num - 1` is not in the set),
     - Initialize a counter and incrementally check for consecutive numbers (`num + 1`, `num + 2`, ...).
     - Track the length of the current sequence.
   - Update the maximum sequence length found.

**Time complexity:** $O(n)$  
**Space complexity:** $O(n)$

## Example

**Input:** nums = [100, 4, 200, 1, 3, 2]  
**Output:** 4  
**Explanation:** The longest consecutive sequence is [1, 2, 3, 4].

**Input:** nums = [0,3,7,2,5,8,4,6,0,1]  
**Output:** 9  
**Explanation:** The longest consecutive sequence is [0,1,2,3,4,5,6,7,8].

## Usage

See `main.go` for code to compute the length of the longest consecutive sequence.

## Resources

- [LeetCode problem](https://leetcode.com/problems/longest-consecutive-sequence/description/)
- [Longest Consecutive Sequence explanation (YouTube)](https://youtu.be/P6RZZMu_maU?si=oPYELqCdIPrsUC7K)
