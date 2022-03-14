// Name: Insert in a Binary Search Tree
// Tags: Binary Search Tree, recursion
// Stats: 16ms-97.70%, 7.2mb-99.65%

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil { return &TreeNode{val, nil, nil} }
	
	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

// Less clean, but slight faster (maybe?) version
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil { return &TreeNode{val, nil, nil} }

	if val < root.Val {
        if root.Left == nil {
            root.Left = &TreeNode{val, nil, nil}
        } else {
		    insertIntoBST(root.Left, val)
        }
	} else {
        if root.Right == nil {
            root.Right = &TreeNode{val, nil, nil}
        } else {
		    insertIntoBST(root.Right, val)
        }
	}
    return root
}


