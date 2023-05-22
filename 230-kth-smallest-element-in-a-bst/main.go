package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func main() {
	fmt.Println("Not able to test it")
	fmt.Println("it works fine, but I failed to pass parameters and I need to figure out how to cast an array to a binary tree")
}

func kthSmallestElement(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0, k)
	// current := root
	visits := 0

	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]		// "pop" the last element from the map

		visits++
		if visits == k {
			return root.Val
		}

		root = root.Right
	}

	return -1
}
