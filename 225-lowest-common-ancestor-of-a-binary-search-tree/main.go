package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println("Not able to test...")
	fmt.Println("It works fine, but I failed to pass parameters and I need to figure out how to cast an array to a binary tree")
}

func lowestCommonAncestorNotWorking(root, p, q *TreeNode) *TreeNode {
	// current := root

	// for current != nil {
	// 	if p.Val > current.Val && q.Val > current.Val {
	// 		current = current.Right
	// 	} else if p.Val < current.Val && q.Val < current.Val {
	// 		current = current.Left
	// 	} else {
	// 		return current
	// 	}
	// }
	return root 	// <- wrong
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} 

	return root
}
