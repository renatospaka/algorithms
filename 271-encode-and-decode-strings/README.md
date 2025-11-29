
# 271. Encode and Decode Strings

[LeetCode problem](https://leetcode.com/problems/encode-and-decode-strings/description/)

## Problem

Design an algorithm to encode a list of strings to a single string and decode it back to the list of strings.

Implement the encode and decode methods:

- `encode(strs)`: Encodes a list of strings to a single string.
- `decode(s)`: Decodes a single string to a list of strings.

## Approach

To avoid ambiguity, encode each string by prefixing it with its length and a special delimiter (such as `#`). For example, the string "hello" becomes "5#hello". Concatenate all encoded strings. To decode, read the length, skip the delimiter, and extract the substring of that length, repeating until the end of the string.

1. For each string, append its length, a delimiter, and the string itself to the result.
2. To decode, iterate through the encoded string, parse the length, extract the string, and repeat.

**Time complexity:** $O(N)$, where $N$ is the total length of all strings.
**Space complexity:** $O(N)$

## Example

**Input:** ["leet","code","!"]  
**Encoded:** "4#leet4#code1#!"  
**Decoded:** ["leet","code","!"]

**Input:** [""]  
**Encoded:** "0#"  
**Decoded:** [""]

## Usage

See `main.go` for code to encode and decode a list of strings.

## Resources

- [LeetCode problem](https://leetcode.com/problems/encode-and-decode-strings/description/)
- [Encode and Decode Strings (YouTube)](https://youtu.be/YPTqKIgVk-k?si=VzEzVYZX501SVjki)
