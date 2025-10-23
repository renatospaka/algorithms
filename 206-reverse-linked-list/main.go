package main

import "fmt"

var (
	list                *ListNode
	inverted            *ListNode
	invertedRecursive   *ListNode
	invertedWithmethods *ListNode
	numbers             []int
)

func main() {
	test([]int{1, 2, 3, 4, 5})
	test([]int{1, 2})
	test([]int{})
	test([]int{4, 6, 77, 43, 10, 32, 89})
}

func test(numbers []int) {
	list = NewListNode()
	list.Push(numbers)
	inverted = reverseList(list)
	fmt.Printf("The inverse of %v is %v\n", numbers, inverted.ToArray())

	list = NewListNode()
	list.Push(numbers)
	invertedRecursive = reverseListRecursive(list)
	fmt.Printf("The inverse of %v is %v using recursivity\n", numbers, invertedRecursive.ToArray())

	list = NewListNode()
	list.Push(numbers)
	invertedWithmethods = reverseListUsingMethods(list)
	fmt.Printf("The inverse of %v is %v using methods\n", numbers, invertedWithmethods.ToArray())
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var previous *ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}
	return previous
}

func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head
	if head.Next != nil {
		newHead = reverseListRecursive(head.Next)
		head.Next.Next = head
	}
	head.Next = nil
	return newHead
}

func reverseListUsingMethods(head *ListNode) *ListNode {
	return head.ToInverse()
}
