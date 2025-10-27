## 279. Perfect Squares

### Problem

Given an integer n, return the least number of perfect square numbers (for example, 1, 4, 9, 16, ...) which sum to n.

For example, given n = 12, the answer is 3 because 12 = 4 + 4 + 4. Given n = 13, the answer is 2 because 13 = 4 + 9.

([LeetCode problem](https://leetcode.com/problems/perfect-squares/))

### Approach

Use dynamic programming to solve the problem:

- Create a dp array of size n+1, where dp[i] represents the least number of perfect squares that sum to i.
- Initialize dp[0] = 0.
- For each i from 1 to n:
    - For each perfect square j*j <= i:
        - dp[i] = min(dp[i], dp[i - j*j] + 1)
- Return dp[n].

Time complexity: $O(n \sqrt{n})$
Space complexity: $O(n)$

### Example

Input: n = 12
Output: 3
Explanation: 12 = 4 + 4 + 4

Input: n = 13
Output: 2
Explanation: 13 = 4 + 9

### Usage

See `main.go` for code to compute the least number of perfect squares that sum to a given number n.

### Resources

- [LeetCode problem](https://leetcode.com/problems/perfect-squares/)
- [Perfect Squares explanation (YouTube)](https://youtu.be/HLZLwjzIVGo?si=ZZ8PtmtCkZh5pBtG)
