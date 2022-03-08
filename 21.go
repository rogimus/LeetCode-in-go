// Name: Merge Two Sorted Lists
// Num: 21
// Tags: Linked Lists, Sorted Lists, Recursion
// Stats: 0ms-100%, 2.6mb-68.77%

package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func soln (list1, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	
	curr1 := list1
	curr2 := list2
	
	if list1.Val < list2.Val {
		l1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		l2.Next = mergeTwoLists(list2.Next, list1)
		return list2
	}
}



func main () {
	input := "."
	fmt.Println(soln(input))
}
