
## 19. Remove Nth Node From End of List

### Problem
Remove the nth node from the end of a singly linked list. ([LeetCode problem](https://leetcode.com/problems/remove-nth-node-from-end-of-list/))

### Approach
This solution removes the nth node from the end in-place with $O(n)$ time and $O(1)$ space.

**Algorithm:**
- Use two pointers (slow and fast) to traverse the list.
- Move the slow pointer to the end, counting steps.
- Move the fast pointer only after the first n steps.
- When slow reaches the end, fast points to the node before the one to remove.
- Adjust pointers to remove the target node.

### Example
Input: `1 -> 2 -> 3 -> 4 -> 5`, n = 2  
Output: `1 -> 2 -> 3 -> 5`

### Usage
See `main.go` for code to create a linked list and remove the nth node from the end.

### Resources
- [LeetCode problem](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)
- [Remove Nth Node explanation (YouTube)](https://youtu.be/XVuQxVej6y8?si=Ew3NQ-zRlobgEh1t)
