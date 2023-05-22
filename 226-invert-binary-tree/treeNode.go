package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode() *TreeNode {
	return &TreeNode{
		Left:  nil,
		Right: nil,
	}
}

// [4, 2, 7, 1, 3, 6, 9]
//					==>
// 							4
//						/		\_
//					2					7
//				/		\			/		\
//			1				3		6			9
//					==>
func (t *TreeNode) Push(numbers []int) {
	if len(numbers) < 2 {
		return
	}
	
	t.Val, t.Left, t.Right = numbers[0], nil, nil
	fmt.Printf("Push -> t.Val: %d | t.Left: %t | t.Right: %t\n", t.Val, t.Left == nil, t.Right == nil)
	for i := 1; i < len(numbers); i++ {
		fmt.Printf("Push -> i: %d => value: %d\n", i, numbers[i])
		t.Add(numbers[i])
		fmt.Println("===================================")
	}
}

func (t *TreeNode) Add(value int) {
	if t == nil {
		// fmt.Println("t == nil")
		return
	}

	// fmt.Printf("Add -> value: %d ! t.Val: %d\n", value, t.Val)
	if value <= t.Val {
		// fmt.Println("value <= t.Val = LEFT")
		if t.Left == nil {
			// fmt.Println("t.Left == nil")
			t.Left = &TreeNode{Val: value, Left: nil, Right: nil}
		} else {
			t.Left.Add(value)
		}
		// fmt.Printf("Left => Val: %d || LeftNull: %t, RightNull: %t\n", t.Left.Val, t.Left.Left == nil,  t.Left.Right == nil)

		} else {
		// fmt.Println("value > t.Val = RIGHT")
		if t.Right == nil {
			// fmt.Println("t.Right == nil")
			t.Right = &TreeNode{Val: value, Left: nil, Right: nil}
		} else {
			t.Right.Add(value)
		}
		// fmt.Printf("Right => Val: %d || LeftNull: %t, RightNull: %t\n", t.Right.Val, t.Right.Left == nil, t.Right.Right == nil)
	}
}

func (t *TreeNode) List() {
	// fmt.Printf("nodes: %v\n", t.ToArray())
}

// [4, 2, 7, 1, 3, 6, 9]
//					==>
// 							4
//						/		\_
//					2					7
//				/		\			/		\
//			1				3		6			9
//					==>
func (t *TreeNode) ToArray() []int {
	treeNode := []int{}
	root := t

	treeNode = append(treeNode, t.Val)
	left, right := getNode(root)
	fmt.Printf("Left => Val: %d || LeftNull: %t, RightNull: %t\n", left.Val, left.Left == nil, left.Right == nil)
	fmt.Printf("Right => Val: %d || LeftNull: %t, RightNull: %t\n", right.Val, right.Left == nil, right.Right == nil)
	for left != nil && right != nil {
		if left != nil {
			treeNode = append(treeNode, left.Val)
		}

		if right != nil {
			treeNode = append(treeNode, right.Val)
		}
	}
	return treeNode
}

func getNode(tree *TreeNode) (*TreeNode, *TreeNode) {
	return tree.Left, tree.Right
}