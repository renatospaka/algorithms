package main

import "fmt"

// type Node struct {
// 	Value int
// 	Next *Node
// }

// type ListNode struct {
// 	Head *Node
// }

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	list1 := []int{2,4,3}
	list2 := []int{2,4,3}
	// need to transform list1 and list2 in ListNode
	// output := addTwoNumbers(list1, list2)
	output := []int{7,0,8}
	fmt.Printf("The sum of %v and %v is %v\n", list1, list2, output)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		value1 := 0
		if l1 != nil {
			value1 = l1.Val
		}

		value2 := 0
		if l2 != nil {
			value2 = l2.Val
		}

		value := value1 + value2 + carry
		carry = value / 10
		value = value % 10
		current.Next = &ListNode{Val: value}

		current = current.Next
		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = nil
		}

		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = nil
		}
	}

	return dummy.Next
}


// func (l *ListNode) insert(value int) {
// 	list := &Node{Value: value, Next: nil}
// 	if l.Head == nil {
// 		l.Head = list
// 	} else {
// 		pointer := l.Head
// 		for pointer.Next != nil {
// 			pointer = pointer.Next
// 		}
// 		pointer.Next = list
// 	}
// }

// func (l *ListNode) show() {
// 	head := l.Head
// 	for head != nil {
// 		fmt.Printf("-> %v", head.Value)
// 		head = head.Next
// 	}
// 	fmt.Println("")
// }
