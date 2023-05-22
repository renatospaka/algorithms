package main

import "fmt"

func main() {
	// fmt.Println("Not able to test it")
	// fmt.Println("it works fine, but I failed to pass parameters and I need to figure out how to cast an array to a binary tree")

	// // ##CASE 1
	// // Input: root = [4,2,7,1,3,6,9]
	// // Output: [4,7,2,9,6,3,1]

	// // ##CASE 2
	// // Input: root = [2,1,3]
	// // Output: [2,3,1]

	// // ##CASE 3
	// // Input: root = [2,1,3]
	// // Output: [2,3,1]

	// // ##CASE 4
	// // Input: root = [100, -20, -50, -15, -60, 50, 60, 55, 85, 15,5, -10]
	// // Output: [100,-50,-20,60,50,-60,-15,null,null,null,-10,5,15,85,55]
	
	var (
		numbers []int
		tree    *TreeNode
		invert  *TreeNode
	)

	numbers = []int{4, 2, 7, 1, 3, 6, 9}
	// numbers = []int{4, 2, 7}
	tree = NewTreeNode()
	tree.Push(numbers)
	fmt.Printf("toArray: %v\n", tree.ToArray())
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
