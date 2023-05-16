package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func isSubtree(root *TreeNode, sub *TreeNode) bool {
	// an empty or null tree can be attached to any final branch of the tree
	if sub == nil {
		return true
	}

	// however, empty or null trees cannot be attached under any circumstancess
	if root == nil {
		return false
	}

	if isSameTree(root, sub) {
		return true
	}
	return isSubtree(root.Left, sub) || isSubtree(root.Right, sub)
}

func isSameTree(tree1 *TreeNode, tree2 *TreeNode) bool {
	if tree1 == nil && tree2 == nil {
		return true
	}

	if tree1 != nil && tree2 != nil && tree1.Val == tree2.Val {
		return isSameTree(tree1.Left, tree2.Left) && isSameTree(tree1.Right, tree2.Right)
	}

	return false
}
