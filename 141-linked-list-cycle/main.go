package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := test1()
	fmt.Printf("is the 1st provided list cickled? %t\n", list1)

	list2 := test2()
	fmt.Printf("is the 2nd provided list cickled? %t\n", list2)

	list3 := test3()
	fmt.Printf("is the 3rd provided list cickled? %t\n", list3)
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

func test1() bool {
	node4 := &ListNode{Val: -4, Next: nil}
	node3 := &ListNode{Val: 0, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	head := &ListNode{Val: 3, Next: node2}
	node4.Next = node2

	return hasCycle(head)
}

func test2() bool {
	node2 := &ListNode{Val: 1, Next: nil}
	head := &ListNode{Val: 2, Next: node2}
	node2.Next = head

	return hasCycle(head)
}

func test3() bool {
	head := &ListNode{Val: 1, Next: nil}

	return hasCycle(head)
}
