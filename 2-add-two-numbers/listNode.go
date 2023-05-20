package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Add(value int) {
	root := l

	for {
		if root.Next == nil {
			root.Next = &ListNode{Val: value, Next: nil}
			break
		}
		root = root.Next
	}
}

func (l *ListNode) List() {
	fmt.Printf("nodes: %v\n", l.ToArray())
}

func (l *ListNode) ToArray() []int {
	listNode := []int{}
	root := l

	for root != nil {
		listNode = append(listNode, root.Val)
		root = root.Next
	}
	return listNode
}
