package main

import (
	"fmt"

	"github.com/renatospaka/linkedlist"
)

var (
	merged *linkedlist.ListNode
	n1     []int
	n2     []int
	n3     []int
	n4     []int
	n5     []int
)

func main() {
	n1 = []int{1, 4, 5}
	n2 = []int{1, 3, 4}
	n3 = []int{2, 6}
	test(n1, n2, n3, n4, n5)

	n1 = nil
	n2 = nil
	n3 = nil
	n4 = nil
	n5 = nil
	test(n1, n2, n3, n4, n5)

	n1 = []int{}
	n2 = nil
	n3 = nil
	n4 = nil
	n5 = nil
	test(n1, n2, n3, n4, n5)
}

func test(n ...[]int) {
	lists := []*linkedlist.ListNode{}
	for _, nums := range n {
		if nums == nil || len(nums) == 0 {
			continue
		}
		list := linkedlist.NewListNode()
		list.Push(nums)
		lists = append(lists, list)
	}

	merged = mergedKLists(lists)
	fmt.Printf("This is the result of merging %v and %v: %v\n", n1, n2, merged.ToArray())
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
func mergedKLists(lists []*linkedlist.ListNode) *linkedlist.ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		mergedLists := []*linkedlist.ListNode{}

		for i := 0; i < len(lists); i += 2 {
			var list1, list2 *linkedlist.ListNode = nil, nil
			// var list2 *linkedlist.ListNode = nil

			list1 = lists[i]
			if i+1 < len(lists) {
				list2 = lists[i+1]
			}
			merged = mergeTwoLists(list1, list2)
			mergedLists = append(mergedLists, merged)
		}
		lists = mergedLists
	}
	return lists[0]
}

func mergeTwoLists(list1, list2 *linkedlist.ListNode) *linkedlist.ListNode {
	temp := &linkedlist.ListNode{}
	current := temp

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// Append any remaining nodes from either list
	if list1 != nil {
		current.Next = list1
	} else if list2 != nil {
		current.Next = list2
	}
	return temp.Next
}
