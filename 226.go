// Name: Invery Binary Tree
// Tags: Binary Tree, recursion
// Stats: 0ms-100%, 2.1mb-100%

func invertTree(root *TreeNode) *TreeNode {
	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
