// Name: Search in a Binary Search Tree
// Tags: Binary Search Tree, recursion
// Stats: 19ms-92.31%, 6.9mb-97.60%

func searchBST(root *TreeNode, val int) *TreeNode {

	if root == nil { return nil }

	if root.Val == val { return root}

	if root.Val > val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}
