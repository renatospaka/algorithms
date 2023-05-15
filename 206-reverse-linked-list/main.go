package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	var (
		listNodes []int
		listNode *ListNode
		reverse *ListNode
	)

	listNodes = []int{1, 2, 3, 4, 5}
	listNode = toListNode(listNodes)
	reverse = reverseList(listNode)
	fmt.Printf("The reverse list of %v is %v\n", listNodes, reverse)
}


func toListNode(numbers []int) *ListNode {
	var (
		current *ListNode
		previous *ListNode
	)

	current = &ListNode{}
	for i := 0; i < len(numbers); i++ {
		current.Val = numbers[i]
		if i == 0 {
			current.Next = nil
		}
		if i > 0 {
			previous.Next = current
		}
		
		previous = current
	}

	// ln := current
	// for ln.Next != nil {
	// 	fmt.Printf("Value: %d, Next: %v\n", current.Val, current.Next)
	// 	ln = ln.Next
	// }
	return current
}


func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var previous *ListNode
	current :=head

	for current != nil {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}
	return previous
}
