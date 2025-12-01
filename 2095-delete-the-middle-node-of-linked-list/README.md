# 2095. Delete the Middle Node of a Linked List

**Difficulty:** Medium

## Problem

You are given the `head` of a linked list. Delete the middle node, and return the `head` of the modified linked list.

The middle node of a linked list of size $n$ is the $\lfloor n / 2 \rfloor$-th node from the start using 0-based indexing, where $\lfloor x \rfloor$ denotes the largest integer less than or equal to $x$.


## Approach

This solution deletes the middle node of a singly linked list in-place with $O(n)$ time and $O(1)$ space.

**Algorithm:**

- Use two pointers, slow and fast, both starting at the head, and a prev pointer to track the node before slow.
- Move slow by one step and fast by two steps in each iteration.
- When fast reaches the end, slow will be at the middle node, and prev will be right before it.
- Remove the middle node by setting `prev.Next = slow.Next`.


## Examples

**Example 1:**

Input: `head = [1,3,4,7,1,2,6]`

Output: `[1,3,4,1,2,6]`

**Example 2:**

Input: `head = [1,2,3,4]`

Output: `[1,2,4]`

**Example 3:**

Input: `head = [2,1]`

Output: `[2]`


## Constraints

- The number of nodes in the list is in the range `[1, 10^5]`.
- `1 <= Node.val <= 10^5`

## Usage

See `main.go` for code to create a linked list and delete the middle node.

## Resources

- [LeetCode problem](https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/description/)
- [Delete the Middle Node of a Linked List explanation (YouTube)](https://www.youtube.com/watch?v=7bAzwFijLJU/)
