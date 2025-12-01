# 876. Middle of the Linked List

**Difficulty:** Easy

## Problem

Given the head of a singly linked list, return the middle node of the linked list. If there are two middle nodes, return the second middle node. 

## Approach

This solution finds the middle node of a singly linked list in-place with $O(n)$ time and $O(1)$ space.

**Algorithm:**

- Use two pointers, slow and fast, both starting at the head.
- Move slow by one step and fast by two steps in each iteration.
- When fast reaches the end, slow will be at the middle node.

## Example

Input: `1 -> 2 -> 3 -> 4 -> 5`  
Output: `3`

Input: `1 -> 2 -> 3 -> 4 -> 5 -> 6`  
Output: `4`

## Usage

See `main.go` for code to create a linked list and find the middle node.

## Resources

- [LeetCode problem](https://leetcode.com/problems/middle-of-the-linked-list/description/)
- [Middle of the Linked List explanation (YouTube)](https://youtu.be/A2_ldqM4QcY?si=O7SuYCNF4d-od9uM)
