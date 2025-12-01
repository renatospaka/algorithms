# 11. Container With Most Water

[LeetCode problem](https://neetcode.io/solutions/longest-consecutive-sequence)

## Problem

Given `n` non-negative integers `height` where each represents a point at coordinate `(i, height[i])`, `n` vertical lines are drawn such that the two endpoints of the line `i` are at `(i, 0)` and `(i, height[i])`. Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

## Approach

Use the two-pointer technique to maximize the area:

1. Initialize two pointers, `left` at the start and `right` at the end of the array.
2. While `left < right`:
	- Calculate the area formed by the lines at `left` and `right`.
	- Update the maximum area if the current area is larger.
	- Move the pointer pointing to the shorter line inward (increment `left` or decrement `right`).

**Time complexity:** $O(n)$  
**Space complexity:** $O(1)$

## Example

**Input:** height = [1,8,6,2,5,4,8,3,7]  
**Output:** 49  
**Explanation:** The lines at index 1 and 8 form a container with area 49.

**Input:** height = [1,1]  
**Output:** 1  
**Explanation:** The lines at index 0 and 1 form a container with area 1.

## Usage

See `main.go` for code to compute the maximum area of water a container can store.

## Resources

- [LeetCode problem](https://neetcode.io/solutions/longest-consecutive-sequence)
- [Container With Most Water (YouTube)](https://youtu.be/YPTqKIgVk-k?si=VzEzVYZX501SVjki)
