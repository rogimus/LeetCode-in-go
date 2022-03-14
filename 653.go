// Name: Validate Binary Search Tree
// Tags: Binary Search Tree, Depth First Search, Preorder Traversal 
// Stats: 16ms-98.87%, 7.4mb-83.62%

// recursion would have been much easier but wanted to practise
// implementation with a queue.
func soln1(root *TreeNode, k int) bool {
	if root == nil { return false }

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	checked := make(map[int]bool, 0)
	
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		
		_,in1 := checked[curr.Val]
		_,in2 := checked[k-curr.Val]
		if in1 != true && in2 != true {
			if k-curr.Val == curr.Val { 
				if curr.Left != nil {
					queue = append(queue, curr.Left)
				}
				if curr.Right != nil {
					queue = append(queue, curr.Right)
				}
				continue 
			}
			t := check(root, k-curr.Val)
			if t { return true }
			checked[curr.Val] = true
		}
		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}
	return false
}

func check(root *TreeNode, val int) bool {
	if root == nil { return false }
	if root.Val == val { return true }
	if val > root.Val { return check(root.Right, val) }
	return check(root.Left, val)
}
