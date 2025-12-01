# 263. Ugly Number

**Difficulty:** Easy

## Problem

An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.

Given an integer n, return true if n is an ugly number.

([LeetCode problem](https://leetcode.com/problems/ugly-number/description/))

## Approach


This solution checks if a number is an ugly number by dividing n by 2, 3, and 5 as long as it is divisible by them, and then checks if the result is 1. The time complexity is $O(\log n)$ and space complexity is $O(1)$.


**Algorithm:**

- If n <= 0, return false.
- While n is divisible by 2, divide n by 2.
- While n is divisible by 3, divide n by 3.
- While n is divisible by 5, divide n by 5.
- If n becomes 1, it is an ugly number; otherwise, it is not.


## Example

Input: n = 6
Output: true
Explanation: 6 = 2 × 3

Input: n = 14
Output: false
Explanation: 14 is not ugly since it includes the prime factor 7.


## Usage

See `main.go` for code to check if a number is an ugly number.

## Resources

- [LeetCode problem](https://leetcode.com/problems/ugly-number/description/)
- [Ugly Number explanation (YouTube)](https://youtu.be/M0Zay1Qr9ws?si=EQ4jlCTNMzPa-xAo)
