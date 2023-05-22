package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println("Not able to test it")
	fmt.Println("it works fine, but I failed to pass parameters and I need to figure out how to cast an array to a binary tree")

	// ##CASE 1
	// Input: root = [2,1,3]
	// Output: true

	// ##CASE 2
	// Input: root = [5,1,4,null,null,3,6]
	// Output: false

	// ##CASE 3
	// Input: root = [5,1,7,3,4,8,9]
	// Output: true
}

func isValidBST(root *TreeNode) bool {	
	return validate(root, nil, nil)
}

func validate(node *TreeNode, left *TreeNode, right *TreeNode) bool {
	if node == nil {
		return true
	}

	if left != nil && node.Val <= left.Val {
		return false
	} 

	if right != nil && node.Val >= right.Val {
		return false
	}

	return (validate(node.Left, left, node) && validate(node.Right, node, right))
}
