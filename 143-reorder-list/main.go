package main

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {

}


func reorderList(head *ListNode) {
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

func reverse(node *ListNode) *ListNode {
	var previous, current *ListNode = nil, node
	for current != nil {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}

	return previous
}
