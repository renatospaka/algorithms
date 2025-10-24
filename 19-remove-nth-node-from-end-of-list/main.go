package main

import (
	"fmt"

	"github.com/renatospaka/linkedlist"
)

var (
	arranged *linkedlist.ListNode
	numbers  []int
	n        int
)

func main() {
	numbers = []int{1, 2, 3, 4, 5}
	n = 2
	test(numbers, n)

	numbers = []int{1}
	n = 1
	test(numbers, n)

	numbers = []int{1, 2}
	n = 1
	test(numbers, n)
}

func test(numbers []int, pos int) {
	list := linkedlist.NewListNode()
	list.Push(numbers)

	arranged = removeNthFromEnd(list, pos)
	fmt.Printf("List with %dth item removed from %v is %v\n", pos, numbers, arranged.ToArray())
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
func removeNthFromEnd(head *linkedlist.ListNode, n int) *linkedlist.ListNode {
	slow, fast, temp, count := head, head, 0, 0
	for slow != nil {
		fmt.Printf("%s | n: %d | ", "S", n)
		printLog(temp, slow, fast)
		if temp > n {
			fmt.Println("temp is greather than n")
			fast = fast.Next
		}
		slow = slow.Next
		temp += 1
		count++
		fmt.Printf("%s | n: %d | ", "E", n)
		printLog(temp, slow, fast)
		// fmt.Println("S | n: 2 | temp: 0 | slow.Val: 1 | fast.Val: 1")
		fmt.Println("-------------------------------------------------")
	}
	fmt.Println("finished loop")

	fmt.Printf("%s | n: %d | ", "B", n)
	printLog(temp, slow, fast)
	if n == temp {
		head = head.Next
	} else {
		fast.Next = fast.Next.Next
	}
	fmt.Println("balanced")
	fmt.Printf("%s | n: %d | ", "F", n)
	printLog(temp, slow, fast)

	return head
}

func printLog(i int, slow, fast *linkedlist.ListNode) {
	fmt.Printf("temp: %d | ", i)
	if slow != nil {
		fmt.Printf("slow.Val: %d | ", slow.Val)
	} else {
		fmt.Printf("slow.Val: nil | ")
	}
	if fast != nil {
		fmt.Printf("fast.Val: %d\n", fast.Val)
	} else {
		fmt.Printf("fast.Val: nil\n")
	}
}
