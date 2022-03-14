// Name: Symmetric Tree
// Number: 101
// Tags: Binary Tree, recursion
// Stats: 0ms-100%, 2.9mb-86.98%

func soln(root *TreeNode) bool {
	if root == nil { return true } 
	return treeIsEqual(root.Left, root.Right)
}


func treeIsEqual (root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil { return true }
	if root1 == nil || root2 == nil { return false }
	if root1.Val != root2.Val {return false }
	
	l := treeIsEqual(root1.Left, root2.Right)
	r := treeIsEqual(root1.Right, root2.Left)		
	if l && r {
		return true
	}
	return false
}
