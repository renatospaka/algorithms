package main

import "fmt"

type Node struct {
	Value int
	Next *Node
}

type ListNode struct {
	Head *Node
}

// type ListNode struct {
// 	Value int
// 	Next *ListNode
// }

func main() {
	list1 := &ListNode{}
	list1.insert(2)
	list1.insert(4)
	list1.insert(3)
	list1.show()
	// listNode1 := list1.toListNode()
	
	list2 := &ListNode{}
	list2.insert(5)
	list2.insert(6)
	list2.insert(4)
	list2.show()
	// listNode2 := list2.toListNode()

	newList := addTwoNumbers(listNode1, listNode2)
	newList.show()
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		head1 = l1.Head
		head2 = l2.Head

		newLN1 = head1
		newLN2 = head2

		head = new(ListNode)
		current = head

		diff int
		sum int
	)
	
	for newLN1 != nil || newLN2 != nil {
		var (
			val1 int
			val2 int
		)
		
		if newLN1 != nil {
			val1 = newLN1.Value
			newLN1 = newLN1.Next
		}
		
		if newLN2 != nil {
			val2 = newLN2.Value
			newLN2 = newLN2.Next
		}

		sum = diff + val1 + val2
		diff = sum / 10
		// // if current.Head == nil {
		// // 	node := &Node{Value: sum % 10}
		// // 	current.Head = node
		// // }
		node := &Node{Value: sum % 10}
		current.Head.Next = node
		current = current.Head.Next
	}

	if diff > 0 {
		current.Next = &ListNode{Value: diff}
	}

	return head.Next
}

func addTwoNumbersOrig(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		newLN1 = l1
		newLN2 = l1

		head = new(ListNode)
		current = head

		diff int
		sum int
	)
	
	for newLN1 != nil || newLN2 != nil {
		var (
			val1 int
			val2 int
		)
		
		if newLN1 != nil {
			val1 = newLN1.Value
			newLN1 = newLN1.Next
		}
		
		if newLN2 != nil {
			val2 = newLN2.Value
			newLN2 = newLN2.Next
		}

		sum = diff + val1 + val2
		diff = sum / 10
		current.Next = &ListNode{Value: sum % 10}
		current = current.Next
	}

	if diff > 0 {
		current.Next = &ListNode{Value: diff}
	}

	return head.Next
}


func (l *ListNode) insert(value int) {
	list := &Node{Value: value, Next: nil}
	if l.Head == nil {
		l.Head = list
	} else {
		pointer := l.Head
		for pointer.Next != nil {
			pointer = pointer.Next
		}
		pointer.Next = list
	}
}

func (l *ListNode) show() {
	head := l.Head
	for head != nil {
		fmt.Printf("-> %v", head.Value)
		head = head.Next
	}
	fmt.Println("")
}

// func (l *ListNode) toListNode() *ListNode {
// 	listNode := &ListNode{}
// 	pointer := l.Head
// 	for pointer.Next != nil {
// 		fmt.Printf("Value: %d   Next: %v\n", pointer.Value, pointer.Next)
// 		listNode.Value = pointer.Value
// 		listNode.Next = pointer.Next.Next
// 		pointer = pointer.Next
// 	}
// 	return listNode
// }

// func (l *ListNode) show() {
// 	fmt.Println("=============================")
// 	for l.Next != nil {
// 		fmt.Printf("-> %v", l.Value)
// 	}
// }
