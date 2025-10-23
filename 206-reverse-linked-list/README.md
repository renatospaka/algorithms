
## 206. Reverse Linked List

### Problem
Reverse a singly linked list. ([LeetCode problem](https://leetcode.com/problems/reverse-linked-list/))

### Approach
This solution iteratively reverses the linked list in $O(n)$ time and $O(1)$ space.

**Algorithm:**
- Initialize three pointers: `prev`, `curr`, and `next`.
- Traverse the list, reversing the `next` pointer of each node.
- At the end, `prev` will point to the new head of the reversed list.

### Example
Original list: `1 -> 2 -> 3 -> 4 -> 5`  
Reversed list: `5 -> 4 -> 3 -> 2 -> 1`

### Usage
See `main.go` for code to create a linked list and reverse it.

### Resources
- [LeetCode problem](https://leetcode.com/problems/reverse-linked-list/)
- [Reverse Linked List explanation (YouTube)](https://youtu.be/G0_I-ZF0S38?si=TL-wDJftDyvi1jBH)
