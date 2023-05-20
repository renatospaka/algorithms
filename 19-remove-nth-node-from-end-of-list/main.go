package main

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast, temp := head, head, 0
	for slow != nil {
		if temp > n {
			fast = fast.Next
		} 
		slow = slow.Next
		temp += 1
	}

	if n == temp {
		head = head.Next
	} else {
		fast.Next = fast.Next.Next
	}

	return head
}
