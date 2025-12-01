# 264. Ugly Number II

**Difficulty:** Medium

## Problem

An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.

Given an integer n, return the nth ugly number.


## Approach

To find the nth ugly number, use a dynamic programming approach with three pointers, each tracking the next multiple of 2, 3, and 5. At each step, select the minimum among these, add it to the list, and advance the corresponding pointer(s).

**Algorithm:**

- Initialize an array `ugly` with the first ugly number 1.
- Set three pointers: `i2`, `i3`, `i5` to 0.
- For each position from 1 to n-1:
	- Compute next multiples: `next2 = ugly[i2] * 2`, `next3 = ugly[i3] * 3`, `next5 = ugly[i5] * 5`.
	- Choose the minimum among them as the next ugly number and append it to the array.
	- Increment the pointer(s) whose multiple was chosen.
- Return the last element in the array.

Time complexity: $O(n)$
Space complexity: $O(n)$


## Example

Input: n = 10
Output: 12
Explanation: The sequence of ugly numbers is [1, 2, 3, 4, 5, 6, 8, 9, 10, 12]. The 10th ugly number is 12.


## Usage

See `main.go` for code to generate and print ugly numbers. Adapt the code to return the nth ugly number as needed.

## Resources

- [LeetCode problem](https://leetcode.com/problems/ugly-number-ii/description/)
- [Ugly Number explanation (YouTube)](https://youtu.be/1pj2a5bmziY?si=YZaVZ6J9iSPdAg7p)
