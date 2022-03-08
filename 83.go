// Name: Remove Duplicates from Sorted List
// Num: 83
// Tags: Linked Lists
// Stats: 0ms-100%, 3mb-85.59%

package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func soln (head *ListNode) *ListNode {
	curr := head
	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}

func soln2 (head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	ptr1 := head 
	ptr2 := head.Next
	for ptr2 != nil {
		for ptr2 != nil && ptr1.Val == ptr2.Val {
			ptr2 = ptr2.Next
			ptr1.Next = ptr2
		} 
		if ptr2 != nil {
			ptr1, ptr2 = ptr1.Next, ptr2.Next 
		}
	}
	return head
}


func main () {
	node3 := ListNode{2, nil}
	node2 := ListNode{1, &node3}
	node1 := ListNode{1, &node2}
	//fmt.Println(node2)
	//fmt.Println(node1)
	fmt.Println(soln(&node1))
	fmt.Println(soln2(&node1))
}

