package main

import "fmt"

func main() {

	var (
		list     *ListNode
		inverted *ListNode
		numbers  []int
	)

	numbers = []int{1, 2, 3, 4, 5}
	list = NewListNode()
	list.Push(numbers)
	inverted = reverseList(list)
	fmt.Printf("The inverse of %v is %v\n", numbers, inverted.ToArray())
	
	numbers = []int{1, 2}
	list = NewListNode()
	list.Push(numbers)
	inverted = reverseList(list)
	fmt.Printf("The inverse of %v is %v\n", numbers, inverted.ToArray())
	
	numbers = []int{}
	list = NewListNode()
	list.Push(numbers)
	inverted = reverseList(list)
	fmt.Printf("The inverse of %v is %v\n", numbers, inverted.ToArray())
}

func reverseList(head *ListNode) *ListNode {
	// var previous *ListNode
	// current := head

	// for current != nil {
	// 	next := current.Next
	// 	current.Next = previous
	// 	previous = current
	// 	current = next
	// }
	// return previous
	return head.ToInverse()
}
