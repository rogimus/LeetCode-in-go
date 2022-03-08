// Name: Binary Tree - Postorder Traversal
// Number: 145
// Tags: Binary Tree, Depth First Search, Recursion, Stacks
// Stats: 0ms-100%, 2mb-57.78%

package main

import "fmt"

type TreeNode struct {
	Val int
	Left *Right
	TreeNode *TreeNode
}

func postorderTraversal (root *TreeNode) []int {
	if root == nil {
		return nil
	}
	
	var res []int
	if root.Left != nil {
		for _, n := range postorderTraversal(root.Left) {
			res = append(res, n)
		}
	}
	if root.Right!= nil {
		for _, n := range postorderTraversal(root.Right) {
			res = append(res, n)
		}
	}
	res = append(res, root.Val)
	return res
}

// dfs solution
// IGNORE THE FACT THAT I CALLED MY STACKS QUEUES.
// My Language Server doesnt support renaming at the moment :((
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
		if len(queue) > 1 {
			queue = queue[:len(queue)-1]
		} else {
			queue = nil
		}

		_,in := visited[curr]
		if in == false {
			visited[curr] = true
		
			queue.append(curr)
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
		} else {
			res = append(res, curr.Val)
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

