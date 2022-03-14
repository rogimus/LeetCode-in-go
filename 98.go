// Name: Validate Binary Search Tree
// Tags: Binary Search Tree, Recursion
// Stats: 3ms-95.17%, 5.1mb-95.75%

const rangeMin = -2147483649
const rangeMax = 2147483649

func isValidBST(root *TreeNode) bool {

	if root == nil { return true }

	if root.Left != nil && root.Left.Val >= root.Val {
		return false
	}
	if root.Right != nil && root.Right.Val <= root.Val {
		return false
	}
	return recursionLeft(root.Left, rangeMin, root.Val) && recursionRight(root.Right, root.Val, rangeMax)
}


func recursion(root *TreeNode, minVal, maxVal int) bool {
	if root == nil { return true }

	if root.Left != nil && (root.Left.Val >= root.Val || root.Left.Val <= minVal) {
		return false
	}
	if root.Right != nil && (root.Right.Val >= maxVal || root.Right.Val <= root.Val) {
		return false
	}
	return recursion(root.Left, minVal, root.Val) && recursion(root.Right, root.Val, maxVal)
}
