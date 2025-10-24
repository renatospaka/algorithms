package main

import (
	"fmt"

	"github.com/renatospaka/linkedlist"
)

var (
	deleted *linkedlist.ListNode
	numbers []int
)

func main() {
	numbers = []int{1, 3, 4, 7, 1, 2, 6}
	test(numbers)

	numbers = []int{1, 2, 3, 4}
	test(numbers)

	numbers = []int{2, 1}
	test(numbers)

	numbers = []int{8}
	test(numbers)
}

func test(numbers []int) {
	list := linkedlist.NewListNode()
	list.Push(numbers)

	deleted = deleteMiddle(list)
	fmt.Printf("This with the middle of %v deleted is %v\n", numbers, deleted.ToArray())
	fmt.Println("===================================================")
	fmt.Println()
	fmt.Println()
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *linkedlist.ListNode) *linkedlist.ListNode {
	if head.Next == nil {
		return nil
	}

	slow, fast := head, head
	var prev *linkedlist.ListNode = nil

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = prev.Next.Next
	return head
}
