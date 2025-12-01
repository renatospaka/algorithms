# 238. Product of Array Except Self

**Difficulty:** Medium

## Problem

Given an integer array `nums`, return an array `answer` such that `answer[i]` is equal to the product of all the elements of `nums` except `nums[i]`.

The solution must be done **without using division** and in $O(n)$ time.

## Approach

To solve this problem without division and in linear time:

1. Create an output array `answer` where `answer[i]` will eventually hold the product of all elements except `nums[i]`.
2. Traverse the array from left to right, for each index `i`, set `answer[i]` to the product of all elements to the left of `i`.
3. Traverse the array from right to left, for each index `i`, multiply `answer[i]` by the product of all elements to the right of `i`.

**Time complexity:** $O(n)$  
**Space complexity:** $O(1)$ (excluding the output array)

## Example

**Input:** nums = [1,2,3,4]  
**Output:** [24,12,8,6]  
**Explanation:**
 - answer[0] = 2 * 3 * 4 = 24
 - answer[1] = 1 * 3 * 4 = 12
 - answer[2] = 1 * 2 * 4 = 8
 - answer[3] = 1 * 2 * 3 = 6

**Input:** nums = [-1,1,0,-3,3]  
**Output:** [0,0,9,0,0]  
**Explanation:**
 - answer[0] = 1 * 0 * -3 * 3 = 0
 - answer[2] = -1 * 1 * -3 * 3 = 9

## Usage

See `main.go` for code to compute the product of array except self.

## Resources

- [LeetCode problem](https://leetcode.com/problems/product-of-array-except-self/description/)
- [Product of Array Except Self (YouTube)](https://youtu.be/bNvIQI2wAjk?si=e4Bj9OMF1gXCqXxc)
