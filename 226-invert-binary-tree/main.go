package main

import "fmt"

func main() {
	var (
		numbers []int
		tree    *TreeNode
		invert  *TreeNode
	)

	numbers = []int{1, 2, 3, 4, 5}
	numbers = []int{1, 2, 3}
	tree = NewTreeNode()
	tree.Push(numbers)
	// fmt.Printf("toArray: %v\n", tree.ToArray())
	// // invert = invertTreeNode(tree)
	fmt.Printf("The inverion list of %v is %v\n", numbers, invert)
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
