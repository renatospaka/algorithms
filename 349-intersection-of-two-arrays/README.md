

# 349. Intersection of Two Arrays

[LeetCode problem](https://leetcode.com/problems/intersection-of-two-arrays/description/)

## Problem

Given two integer arrays `nums1` and `nums2`, return an array of their intersection. Each element in the result must be unique and you may return the result in any order.

## Approach

Use a set to efficiently find the intersection:

1. Convert one of the arrays (e.g., `nums1`) to a set for $O(1)$ lookups.
2. Iterate through the other array (`nums2`), and for each element, check if it exists in the set.
3. If it does, add it to the result set (to ensure uniqueness).
4. Return the result as a list or array.

**Time complexity:** $O(n + m)$, where $n$ and $m$ are the lengths of the two arrays.  
**Space complexity:** $O(n)$ for the set.

## Example

**Input:** nums1 = [1,2,2,1], nums2 = [2,2]  
**Output:** [2]

**Input:** nums1 = [4,9,5], nums2 = [9,4,9,8,4]  
**Output:** [9,4]  
**Explanation:** [4,9] is also accepted.

## Usage

See `main.go` for code to compute the intersection of two arrays.

## Resources

- [LeetCode problem](https://leetcode.com/problems/intersection-of-two-arrays/description/)
- [Intersection of Two Arrays (YouTube)](https://youtu.be/fwUTXaMom6U?si=-5y-1Vc6GhCKgVCJ)
