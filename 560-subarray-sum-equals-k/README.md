

# 560. Subarray Sum Equals K

[LeetCode problem](https://leetcode.com/problems/subarray-sum-equals-k/description/)

## Problem

Given an array of integers `nums` and an integer `k`, return the total number of continuous subarrays whose sum equals to `k`.

## Approach

Use a hash map to store the cumulative sum and its frequency while iterating through the array:

1. Initialize a hash map with `{0: 1}` to handle the case where a subarray starts from index 0.
2. Iterate through the array, maintaining a running sum.
3. For each element, check if `running_sum - k` exists in the hash map. If it does, add its frequency to the result.
4. Update the hash map with the current running sum.

**Time complexity:** $O(n)$  
**Space complexity:** $O(n)$

## Example

**Input:** nums = [1,1,1], k = 2  
**Output:** 2  
**Explanation:** There are two subarrays [1,1] that sum to 2.

**Input:** nums = [1,2,3], k = 3  
**Output:** 2  
**Explanation:** There are two subarrays: [1,2] and [3].

## Usage

See `main.go` for code to compute the number of subarrays whose sum equals k.

## Resources

- [LeetCode problem](https://leetcode.com/problems/subarray-sum-equals-k/description/)
- [Subarray Sum Equals K explanation (YouTube)](https://youtu.be/fFVZt-6sgyo?si=vjuCdw47LwxohZa4)
