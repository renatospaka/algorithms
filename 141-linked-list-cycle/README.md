## 141. Linked List Cycle

### Problem
Detect if a linked list contains a cycle. ([LeetCode problem](https://leetcode.com/problems/linked-list-cycle/))

### Approach
This solution uses Floyd's Tortoise and Hare algorithm, which detects cycles in $O(n)$ time and $O(1)$ space.

**Algorithm:**
- Use two pointers (`slow` and `fast`).
- Move `slow` by one step and `fast` by two steps.
- If they meet, a cycle exists. If `fast` reaches the end, no cycle exists.

### Example
Linked list: `3 -> 2 -> 0 -> -4`
Cycle: `-4` points back to `2`

### Usage
See `main.go` for code to create a cyclic linked list and detect the cycle.

### Resources
- [LeetCode problem: ](https://leetcode.com/problems/linked-list-cycle/)
- [Floyd's Tortoise and Hare explanation (YouTube)](https://youtu.be/gBTe7lFR3vc?si=LxM8Nv82Ja9XyO8f)
