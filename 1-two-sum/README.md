## 1. Two Sum

### Problem

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

([LeetCode problem](https://leetcode.com/problems/two-sum/))

### Approach

Use a hash map to store the difference between the target and each element:

- Iterate through the array.
- For each element, check if it exists in the hash map.
	- If it does, return the indices.
	- If not, store the difference (target - current element) and its index in the hash map.

Time complexity: $O(n)$  
Space complexity: $O(n)$

### Example

Input: nums = [2,7,11,15], target = 9  
Output: [0,1]  
Explanation: nums[0] + nums[1] = 2 + 7 = 9

Input: nums = [3,2,4], target = 6  
Output: [1,2]

### Usage

See `main.go` for code to compute the indices of the two numbers that add up to the target.

### Resources

- [LeetCode problem](https://leetcode.com/problems/two-sum/)
- [Two Sum explanation (YouTube)](https://youtu.be/KLlXCFG5TnA?si=iYqrignkMcZbOOwm)
