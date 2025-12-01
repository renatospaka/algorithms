# 21. Merge Two Sorted Lists

**Difficulty:** Easy

## Problem

Merge two sorted linked lists and return it as a new sorted list.

## Approach

This solution iteratively merges two sorted linked lists in $O(n)$ time and $O(1)$ space.

**Algorithm:**
- Initialize a dummy node and a tail pointer.
- Traverse both lists, always attaching the smaller node to the merged list.
- When one list is exhausted, attach the remainder of the other list.
- Return the merged list starting from the next of the dummy node.

## Example

List 1: `1 -> 2 -> 4`
List 2: `1 -> 3 -> 4`
Merged list: `1 -> 1 -> 2 -> 3 -> 4 -> 4`

## Usage

See `main.go` for code to create linked lists and merge them.

## Resources

- [LeetCode problem](https://leetcode.com/problems/merge-two-sorted-lists/)
- [Merge Two Sorted Lists explanation (YouTube)](https://youtu.be/XIdigk956u0?si=sni8g-vx7xptCsLw)
