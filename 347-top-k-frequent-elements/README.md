
# 347. Top K Frequent Elements

[LeetCode problem](https://leetcode.com/problems/top-k-frequent-elements/)

## Problem

Given an integer array `nums` and an integer `k`, return the `k` most frequent elements. You may return the answer in any order.

## Approach

Use a hash map to count the frequency of each element, then use a min-heap (priority queue) to keep track of the top k frequent elements:

1. Count the frequency of each element in `nums` using a hash map.
2. Use a min-heap to keep the k elements with the highest frequency.
3. Extract the elements from the heap to get the result.

**Time complexity:** $O(n \log k)$  
**Space complexity:** $O(n)$

## Example

**Input:** nums = [1,1,1,2,2,3], k = 2  
**Output:** [1,2]  
**Explanation:** The two most frequent elements are 1 and 2.

**Input:** nums = [1], k = 1  
**Output:** [1]  
**Explanation:** The most frequent element is 1.

## Usage

See `main.go` for code to compute the k most frequent elements in an array.

## Resources

- [LeetCode problem](https://leetcode.com/problems/top-k-frequent-elements/)
- [Top K Frequent Elements (YouTube)](https://youtu.be/YPTqKIgVk-k?si=VzEzVYZX501SVjki)
