// Name: Binary Tree - Maximum Depth
// Tags: Recursion, Binary Tree
// Stats: 0ms-100%, 4.1mb-89.36%


func recursion (root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftLength := recursion(root.Left)
	rightLength := recursion(root.Right)

	if leftLength > rightLength {
		return leftLength +1
	}
	return rightLength + 1
}
