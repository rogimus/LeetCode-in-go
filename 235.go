// Name: Lowest Common Ancestor of a Binary Search Tree 
// Tags: Binary Search Depth, 
// Stats: 16ms-93.05%, 7.1mb-99.58%

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	minVal := min(p.Val, q.Val)
	maxVal := max(p.Val, q.Val)
	if root.Val >= minVal && root.Val <= maxVal { return root }
	if root.Val < minVal { return lowestCommonAncestor(root.Right, p,q) }
	return lowestCommonAncestor(root.Left, p,q)	
}

func min(p, q int) int {
    if p < q { return p }
    return q
}

func max(p, q int) int {
    if p < q { return q }
    return p
}
