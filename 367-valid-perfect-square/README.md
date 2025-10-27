## 367. Valid Perfect Square

### Problem

Given a positive integer num, return true if num is a perfect square, otherwise return false.

A perfect square is an integer that is the square of an integer; in other words, it is the product of some integer with itself.

([LeetCode problem](https://leetcode.com/problems/valid-perfect-square/description/))

### Approach

To determine if a number is a perfect square without using built-in square root functions, use a binary search approach:

- Set left = 1 and right = num.
- While left <= right:
	- Compute mid = (left + right) / 2.
	- If mid * mid == num, return true.
	- If mid * mid < num, set left = mid + 1.
	- If mid * mid > num, set right = mid - 1.
- If no integer square is found, return false.

Time complexity: $O(\log n)$
Space complexity: $O(1)$

### Example

Input: num = 16
Output: true
Explanation: 4 * 4 = 16, so 16 is a perfect square.

Input: num = 14
Output: false
Explanation: 14 is not a perfect square.

### Usage

See `main.go` for code to check if a number is a perfect square. Adapt the code to return a boolean as needed.

### Resources

- [LeetCode problem](https://leetcode.com/problems/valid-perfect-square/description/)
- [Valid Perfect Square explanation (YouTube)](https://youtu.be/1pj2a5bmziY?si=YZaVZ6J9iSPdAg7p)
