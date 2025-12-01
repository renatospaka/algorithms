# 15. 3Sum

**Difficulty:** Medium

## Problem

Given an integer array `nums`, return all the triplets `[nums[i], nums[j], nums[k]]` such that `i != j`, `i != k`, and `j != k`, and `nums[i] + nums[j] + nums[k] == 0`.

You must not return duplicate triplets.

## Approach

Use sorting and the two-pointer technique:

- Sort the array `nums`.
- Iterate through the array:
  - For each index `i`, skip duplicates.
  - Use two pointers (`left` and `right`) to find pairs such that `nums[i] + nums[left] + nums[right] == 0`.
  - Skip duplicates for `left` and `right`.
- Collect all unique triplets found.

Time complexity: $O(n^2)$  
Space complexity: $O(1)$ (excluding the space for the output)

## Example

Input: nums = [-1,0,1,2,-1,-4]  
Output: [[-1,-1,2],[-1,0,1]]  
Explanation: 
- (-1) + 0 + 1 = 0
- (-1) + (-1) + 2 = 0
The distinct triplets are [[-1,-1,2],[-1,0,1]].

Input: nums = [0,1,1]  
Output: []

Input: nums = [0,0,0]  
Output: [[0,0,0]]

## Usage

See `main.go` for code to compute all unique triplets in an array that sum to zero.

## Resources

- [LeetCode problem](https://leetcode.com/problems/3sum/)
- [3Sum explanation (YouTube)](https://youtu.be/jzZsG8n2R9A?si=ancU237cA-fG2pE7)
