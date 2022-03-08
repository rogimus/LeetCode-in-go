// Name: Reverse Linked List
// Num: 206
// Tags: Linked List
// Stats: 0ms-100%, 2.5mb-100%

package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func soln (head *ListNode) *ListNode {	
	curr := head
	var next *ListNode
	var prev *ListNode = nil
	
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return curr
}


func main () {
	input := "."
	fmt.Println(soln(input))
}

