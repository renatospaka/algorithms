package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

// Recursive Depth First Search DFS
func maxDepthRecursiveDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// left := maxDepthRecursiveDFS(root.Left)
	// right := maxDepthRecursiveDFS(root.Right)
	// if left > right {
	// 	return 1 + left
	// }
	// return 1 + right
	return 1 + int(math.Max(float64(maxDepthRecursiveDFS(root.Left)), float64(maxDepthRecursiveDFS(root.Right))))
}
