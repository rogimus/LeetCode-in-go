// Name: Linked List Cycle
// Tags: Slow/Fast algorithm, Floyd's algorithm
// Stats: 4ms-98.14%, 4.4mb-88.51%

package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func soln (head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head
	for fast != nil && fast.Next !=nil {
		slow = slow.Next
		fast  = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}


func main () {
	var linkedList [4]ListNode
	linkedList[0].Val = 3
	linkedList[0].Next = &linkedList[1]
	linkedList[1].Val = 2
	linkedList[1].Next = &linkedList[2]
	linkedList[2].Val = 0
	linkedList[2].Next = &linkedList[3]
	linkedList[3].Val = -4
	linkedList[3].Next = &linkedList[1]
	head := &linkedList[0]
	fmt.Println(soln(head))
}

