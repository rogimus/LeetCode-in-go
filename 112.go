// Name: Path Sum
// Number: 226
// Tags: Binary Tree, recursion, Depth First Search
// Stats: 0ms-100%, 4.5mb-94.84%

package hasPathSum


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil { return false }
	if root.Val == targetSum && (root.Left == nil && root.Right == nil) { return true }

	l := hasPathSum(root.Left, targetSum - root.Val)
	if l { return l }
	r := hasPathSum(root.Right, targetSum - root.Val)
	if r { return r }

	return false
}
