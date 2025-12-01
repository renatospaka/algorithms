# 167. Two Sum II - Input Array Is Sorted

**Difficulty:** Medium

## Problem

Given a 1-indexed array of integers `numbers` that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number. Return the indices of the two numbers (1-indexed) as an integer array `answer` of size 2, where `1 <= answer[0] < answer[1] <= numbers.length`. You may not use the same element twice.

## Approach

Use the two-pointer technique to solve the problem:

1. Initialize two pointers, `left` at the start and `right` at the end of the array.
2. While `left < right`:
	- Calculate the sum of `numbers[left] + numbers[right]`.
	- If the sum equals the target, return `[left + 1, right + 1]`.
	- If the sum is less than the target, increment `left`.
	- If the sum is greater than the target, decrement `right`.

**Time complexity:** $O(n)$  
**Space complexity:** $O(1)$

## Example

**Input:** numbers = [2,7,11,15], target = 9  
**Output:** [1,2]  
**Explanation:** The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2.

**Input:** numbers = [2,3,4], target = 6  
**Output:** [1,3]  
**Explanation:** The sum of 2 and 4 is 6. Therefore, index1 = 1, index2 = 3.

## Usage

See `main.go` for code to compute the indices of the two numbers that add up to the target.

## Resources

- [LeetCode problem](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/)
- [Two Sum II explanation (YouTube)](https://youtu.be/cQ1Oz4ckceM?si=Zcn0jdmj7sFXEyOy)
