
# 121. Best Time to Buy and Sell Stock

**Difficulty:** Easy

## Problem

You are given an array `prices` where `prices[i]` is the price of a given stock on the $i$-th day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

## Approach

Track the minimum price seen so far and the maximum profit:

- Iterate through the array of prices.
- For each price, update the minimum price if the current price is lower.
- Calculate the profit if selling at the current price and update the maximum profit if this profit is higher.

Time complexity: $O(n)$  
Space complexity: $O(1)$

## Example

Input: prices = [7,1,5,3,6,4]  
Output: 5  
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.

Input: prices = [7,6,4,3,1]  
Output: 0  
Explanation: No transaction is done, i.e. max profit = 0.

## Usage

See `main.go` for code to compute the maximum profit from a list of stock prices.

## Resources

- [LeetCode problem](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/)
- [Best Time to Buy and Sell Stock (YouTube)](https://youtu.be/1pkOgXD63yU?si=wHpgjfcK-hb7vgxV)
