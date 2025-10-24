package main

import (
	"fmt"

	"github.com/renatospaka/linkedlist"
)

var (
	// deleted *linkedlist.ListNode
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

	numbers = []int{0, 12, 3, 7, 2, 1, 5, 0, 4, 8}
	test(numbers)
}

func test(numbers []int) {
	list := linkedlist.NewListNode()
	list.Push(numbers)
	deleted := deleteMiddle(list)
	fmt.Printf("This with the middle of %v deleted is %v\n", numbers, deleted.ToArray())

	list2 := linkedlist.NewListNode()
	list2.Push(numbers)
	deletedFaster := deleteMiddleFaster(list2)
	fmt.Printf("This with the middle of %v deleted is %v\n", numbers, deletedFaster.ToArray())

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

func deleteMiddleFaster(head *linkedlist.ListNode) *linkedlist.ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	var slow, fast *linkedlist.ListNode = nil, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next

		if slow == nil {
			slow = head
		} else {
			slow = slow.Next
		}
	}
	slow.Next = slow.Next.Next
	return head
}
