package main

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {

}

func mergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	temp := new(ListNode)
	tail := temp

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}

		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
	} else if list2 != nil {
		tail.Next = list2
	}

	return temp.Next
}

func mergeTwoList2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil || list2 == nil {
		return &ListNode{}
	}

	pointer1, pointer2 := list1, list2
	processed := new(ListNode)
	temp := processed

	for pointer1 != nil && pointer2 != nil {
		if pointer1.Val < pointer2.Val {
			temp.Next = pointer1
			temp = temp.Next
			pointer1 = pointer1.Next
		} else {
			temp.Next = pointer2
			temp = temp.Next
			pointer2 = pointer2.Next
		}
	}

	for pointer1 != nil {
		temp.Next = pointer1
		temp = temp.Next
		pointer1 = pointer1.Next
	}

	for pointer2 != nil {
		temp.Next = pointer2
		temp = temp.Next
		pointer2 = pointer2.Next
	}

	return temp.Next
}
