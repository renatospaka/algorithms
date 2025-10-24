package main

import (
	"fmt"

	"github.com/renatospaka/linkedlist"
)

var (
	middle  *linkedlist.ListNode
	numbers []int
)

func main() {
	numbers = []int{1, 2, 3, 4, 5}
	test(numbers)

	numbers = []int{1, 2, 3, 4, 5, 6}
	test(numbers)

	numbers = []int{1, 2, 3, 4, 5, 6, 7, 8}
	test(numbers)
}

func test(numbers []int) {
	list := linkedlist.NewListNode()
	list.Push(numbers)

	middle = middleNode(list)
	fmt.Printf("This is the middle of %v: %v\n", numbers, middle.ToArray())
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
func middleNode(head *linkedlist.ListNode) *linkedlist.ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		// printLog(i, slow, fast)
	}

	return slow
}

func printLog(i int, slow, fast *linkedlist.ListNode) {
	fmt.Printf("temp: %d | ", i)
	if slow != nil {
		fmt.Printf("slow.Val: %d | ", slow.Val)
	} else {
		fmt.Printf("slow.Val: nil | ")
	}
	if fast != nil {
		fmt.Printf("fast.Val: %d\n", fast.Val)
	} else {
		fmt.Printf("fast.Val: nil\n")
	}
}
