// Name: Binary Tree - Level Order Traversal
// Tags: Binary Tree, Breadth First Search, Depth First Search, Recursion, Stacks, Queue
// Stats: 0ms-100%, 2.8mb-89.95%

// bfs and recursion are okay. The dfs is so not clean lol.
// bfs was also easily the best memory wise

package p102


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// recursion
func p102(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    
	var res [][]int
    res = make([][]int, 1)

	res[0] = append(res[0], root.Val)
    res = *addLevel(root.Left, 1, &res)
	res = *addLevel(root.Right, 1, &res)
	return res
}

func addLevel(root *TreeNode, level int, ptrRes *[][]int) *[][]int{

	if root == nil {
		return ptrRes
	}

	res := *ptrRes
    
    if len(res) <= level {
        lev := []int{root.Val}
        res = append(res, lev)
    } else {
        	res[level] = append(res[level], root.Val)
    }
	
	res = *addLevel(root.Left, level+1, &res)
    res = *addLevel(root.Right, level+1, &res)
    return &res
}


// dfs.
// visited tracks whether a it's visited.
// stack is the stack for dfs.
// heights tracks the heigh of the node.
// res is the result.
// this is such a mess....
func dfs (root *TreeNode) [][]int {

	var stack []*TreeNode
	stack = make([]*TreeNode, 0)
	stack = append(stack, root)

 	var visited map[*TreeNode]bool
	var heights map[*TreeNode]int
	visited = make(map[*TreeNode]bool, 0)
	heights = make(map[*TreeNode]int, 0)
	heights[root] = 0

	var res [][]int
	res = make([][]int, 2000)

	for len(stack) > 0 {
		currNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if currNode == nil { continue }

		_,in := visited[currNode]
		if in == false {
			visited[currNode] = true

			currHeight,_ := heights[currNode]
			if currNode.Right != nil {
				stack = append(stack, currNode.Right)
				heights[currNode.Right] = currHeight + 1
			}
			if currNode.Left != nil {
				stack = append(stack, currNode.Left)
				heights[currNode.Left] = currHeight + 1
				res[currHeight+1] = append(res[currHeight+1], currNode.Left.Val)
			}
			if currNode.Right != nil {
				res[currHeight+1] = append(res[currHeight+1], currNode.Right.Val)
			}
			
		} else {
			continue
		}
	}
	for i,_ := range res {
		if len(res[i]) == 0 {
			return res[:i]
		}
	}
	return res
}

// bfs
func bfs(root *TreeNode) [][]int {
	if root == nil { return nil }
	
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	res := make([][]int, 0)

	countCurr := 1
	countNext := 1
	
	for len(queue) > 0 {
		currNode := queue[0]
		queue = queue[1:]

		if len(res[len(res)-1]) < countCurr {
			res[len(res)-1] = append(res[len(res)-1], currNode.Val)
		} else {
			countCurr = countNext
			res = append(res, []int{currNode.Val})
		}
		if currNode.Left != nil {
			queue = append(queue, currNode.Left)
			countNext ++
		}
		if currNode.Right != nil {
			queue = append(queue, currNode.Right)
			countNext ++
		}
		countNext -- 
	}
	return res
}
