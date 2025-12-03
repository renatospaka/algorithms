# 70. Climbing Stairs

**Difficulty:** Easy

## Problem

You are climbing a staircase. It takes $n$ steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

## Approach

This is a dynamic programming problem similar to the Fibonacci sequence:

- Use two variables to store the number of ways to reach the previous two steps.
- Iterate from step 3 to $n$, updating the number of ways at each step.
- The total ways to reach the $n$th step is the sum of ways to reach the $(n-1)$th and $(n-2)$th steps.

Time complexity: $O(n)$  
Space complexity: $O(1)$

## Example

Input: n = 2  
Output: 2  
Explanation: There are two ways to climb to the top:
- 1 step + 1 step
- 2 steps

Input: n = 3  
Output: 3  
Explanation: There are three ways to climb to the top:
- 1 step + 1 step + 1 step
- 1 step + 2 steps
- 2 steps + 1 step

## Usage

See `main.go` for code to compute the number of distinct ways to climb a staircase of $n$ steps.

## Resources

- [LeetCode problem](https://leetcode.com/problems/climbing-stairs/)
- [Climbing Stairs explanation (YouTube)](https://youtu.be/Y0lT9Fck7qI?si=ukl1UC3qjZpWTX9l)
