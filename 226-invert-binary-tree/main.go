package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	var (
		treeNodes []int
		treeNode *TreeNode
		reverse *TreeNode
	)

	treeNodes = []int{1, 2, 3, 4, 5}
	treeNode = new(TreeNode)
	reverse = invertTreeNode(treeNode)
	fmt.Printf("The reverse list of %v is %v\n", treeNodes, reverse)
}

func invertTreeNode(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	// now process both left and right branches
	invertTreeNode(root.Left)
	invertTreeNode(root.Right)

	return root
}
