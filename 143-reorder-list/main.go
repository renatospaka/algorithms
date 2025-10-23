package main

import (
	"fmt"

	"github.com/renatospaka/linkedlist"
)

var (
	// list    *linkedlist.ListNode
	// ordered *linkedlist.ListNode
	numbers []int
)

func main() {
	numbers = []int{1, 2, 3, 4}
	test(numbers)

	numbers = []int{1, 2, 3, 4, 5}
	test(numbers)

	numbers = []int{1, 2, 3, 4, 5, 6}
	test(numbers)

	numbers = []int{1, 2, 3, 4, 5, 6, 7}
	test(numbers)
}

func test(numbers []int) {
	list := linkedlist.NewListNode()
	list.Push(numbers)

	reorderList(list)
	fmt.Printf("Ordered of %v is %v\n", numbers, list.ToArray())
	fmt.Println("======================================")
	fmt.Println()
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *linkedlist.ListNode) {
	// middle of the list
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	reversed := reverse(slow.Next)
	slow.Next = nil
	current := head

	for current != nil && reversed != nil {
		next := current.Next
		reversedNext := reversed.Next
		current.Next = reversed
		reversed.Next = next
		current = next
		reversed = reversedNext
	}
}

func reverse(node *linkedlist.ListNode) *linkedlist.ListNode {
	var previous, current *linkedlist.ListNode = nil, node
	for current != nil {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}

	return previous
}
