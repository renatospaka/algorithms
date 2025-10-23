package main

import (
	"fmt"

	"github.com/renatospaka/linkedlist"
)

var (
	firstList  *linkedlist.ListNode
	secondList *linkedlist.ListNode
	merged     *linkedlist.ListNode
	n1, n2     []int
)

func main() {
	n1 = []int{1, 2, 4}
	n2 = []int{1, 3, 4}
	test(n1, n2)

	n1 = nil
	n2 = []int{1, 2}
	test(n1, n2)

	n1 = nil
	n2 = nil
	test(n1, n2)
}

func test(numbers1, numbers2 []int) {
	if numbers1 == nil || len(numbers1) == 0 {
		firstList = nil
	} else {
		firstList = linkedlist.NewListNode()
		firstList.Push(numbers1)
	}

	if numbers2 == nil || len(numbers2) == 0 {
		secondList = nil
	} else {
		secondList = linkedlist.NewListNode()
		secondList.Push(numbers2)
	}

	merged = mergeTwoList(firstList, secondList)
	fmt.Printf("Merge of %v and %v is %v\n", numbers1, numbers2, merged.ToArray())
	fmt.Println("======================================")
	fmt.Println()
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoList(list1 *linkedlist.ListNode, list2 *linkedlist.ListNode) *linkedlist.ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	temp := new(linkedlist.ListNode)
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

func printLog(id int, antesDepois string, list *linkedlist.ListNode) {
	if list != nil {
		if list.Next != nil {
			fmt.Printf("%s=> list%d.Val: %d | list%d.Next.Val: %d\n", antesDepois, id, list.Val, id, list.Next.Val)
		} else {
			fmt.Printf("%s=> list%d.Val: %d | list%d.Next is nil\n", antesDepois, id, list.Val, id)
		}
	} else {
		fmt.Printf("%s=> list%d is nil\n", antesDepois, id)
	}
}
