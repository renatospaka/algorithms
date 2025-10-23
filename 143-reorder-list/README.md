## 143. Reorder List

### Problem
Reorder a linked list so that the nodes are arranged in the following order: first node, last node, second node, second last node, and so on. ([LeetCode problem](https://leetcode.com/problems/reorder-list/))

### Approach
This solution reorders the linked list in-place with $O(n)$ time and $O(1)$ space.

**Algorithm:**
- Find the middle of the linked list using the slow and fast pointer technique.
- Reverse the second half of the list.
- Merge the two halves by alternating nodes from each half.

### Example
Input: `1 -> 2 -> 3 -> 4`
Output: `1 -> 4 -> 2 -> 3`

### Usage
See `main.go` for code to create a linked list and reorder it.

### Resources
- [LeetCode problem](https://leetcode.com/problems/reorder-list/)
- [Reorder List explanation (YouTube)](https://youtu.be/S5bfdUTrKLM?si=_1pXRxq-pCTd5jbx)