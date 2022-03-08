// Name: Binary Tree - Preorder Traversal
// Number: 144
// Tags: Binary Tree, Depth First Search, Recursion
// Stats: 5ms-100%, 2.1mb-61.99%

package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func preorderTraversal (root *TreeNode) []int {
	if root == nil {
		return nil
	}
	
	var res []int
	res = append(res, root.Val)
	if root.Left != nil {
		for _, n := range preorderTraversal(root.Left) {
			res = append(res, n)
		}
	}
	if root.Right != nil {
		for _, n := range preorderTraversal(root.Right) {
			res = append(res, n)
		}
	}
	return res
}

// dfs solution
func dfs (root *TreeNode) []int {
	if root == nil {
		return nil
	}
	
	var res []int
	var queue []*TreeNode
	visited := make(map[*TreeNode]bool)
	var curr *TreeNode

	queue = append(queue, root)

	for queue != nil {
		curr = queue[len(queue)-1]
		res = append(res, curr.Val)
		if len(queue) > 1 {
			queue = queue[:len(queue)-1]
		} else {
			queue = nil
		}
		_,in := visited[curr]
		if in == false {
			visited[curr] = true
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
	}	
	return res
}

func main () {
	n3 := TreeNode{2, nil, nil}
	n2 := TreeNode{1, &n3, nil}
	n1 := TreeNode{3, &n2, nil}
	fmt.Println(dfs(&n1))
}

