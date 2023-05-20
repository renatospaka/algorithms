package main

import "fmt"

func main() {
	list1 := createListNode([]int{2, 4, 3})
	list2 := createListNode([]int{5, 6, 4})
	output := addTwoNumbers(list1, list2)
	fmt.Printf("The sum of %v and %v is %v\n", list1.ToArray(), list2.ToArray(), output.ToArray())
}

func createListNode(numbers []int) *ListNode {
	root := &ListNode{Val: numbers[0], Next: nil}
	for i := 1; i < len(numbers); i++ {
		root.Add(numbers[i])
	}
	return root
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		value1 := 0
		if l1 != nil {
			value1 = l1.Val
		}

		value2 := 0
		if l2 != nil {
			value2 = l2.Val
		}

		value := value1 + value2 + carry
		carry = value / 10
		value = value % 10
		current.Next = &ListNode{Val: value}

		current = current.Next
		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = nil
		}

		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = nil
		}
	}

	return dummy.Next
}
