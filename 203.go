// Name: Remove Linked List Elements
// Num: 203
// Tags: Linked List
// Stats: 3ms-98.31%, 4.7mb-84.82%


package main

import "fmt"

func soln () {
	if head == nil {
		return nil
	}

	t := head
	head = *ListNode{0, t}

	curr := head
	for curr != nil && curr.Val == val {
		curr = curr.Next
	}
	head = curr
	if  head == nil {
		return nil
	}
	
	for curr.Next != nil {
		if curr.Next.Val == val {``
            curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}

func soln2 () {	
	t := ListNode{0, head}
	curr := &t
	
	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return t.Next
}

func main () {
	input := "."
	fmt.Println(soln(input))
}
