# 23. Merge K Sorted Lists

**Difficulty:** Hard

## Problem

You are given an array of `k` linked-lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.  

## Approach

This solution merges k sorted linked lists efficiently using a min-heap (priority queue) with $O(N \log k)$ time and $O(k)$ space, where $N$ is the total number of nodes.

**Algorithm:**

- Initialize a min-heap and insert the head of each non-empty list.
- While the heap is not empty:
  - Extract the node with the smallest value.
  - Add it to the merged list.
  - If the extracted node has a next node, insert it into the heap.
- Continue until all nodes are merged.

## Example

Input:  
`lists = [1->4->5, 1->3->4, 2->6]`  
Output:  
`1->1->2->3->4->4->5->6`

## Usage

See `main.go` for code to create linked lists and merge them.

## Resources

- [LeetCode problem](https://leetcode.com/problems/merge-k-sorted-lists/description/)
- [Merge K Sorted Lists explanation (YouTube)](https://www.youtube.com/watch?v=q5a5OiGbT6Q)
