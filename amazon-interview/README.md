## 1. Two Sum

### Problem

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.


([LeetCode problem](https://leetcode.com/problems/two-sum/))

[See detailed solution in another project](../1-two-sum/README.md)



---

## 2. Complementary Angles (amazon Interview)

### Problem

Given an array of integer angles, print every possible angle pair that form a complement (sum is 90).

- The same pair (indexes) can only be accounted once.
- Example input: `[10, 40, 89, 23, 1, 50, 67, 80, 40, 50]`

### Approach

**Brute-force (O(n²))**
- Use two nested loops to check all unique pairs (i, j) with i < j.
- Print the pair if their sum is 90.

**Optimized (O(n))**
- Use a hash map to store seen values and their indices.
- For each angle, check if (90 - angle) has been seen before.
- Print all unique pairs found this way.

### Example

Input: `[10, 40, 89, 23, 1, 50, 67, 80, 40, 50]`

Output:
```
Pair: (0, 7) => (10, 80)
Pair: (1, 4) => (40, 50)
Pair: (1, 9) => (40, 50)
Pair: (3, 6) => (23, 67)
```

### Usage

See `main.go` for code to print all unique pairs of angles that sum to 90.
