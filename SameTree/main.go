package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return Traverse(p, q)
}

func Traverse(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	} else if p != q {
		return false
	}

	return true
}

//[2,1,3,7,4,9,5,8,6]
func main() {
	root1 := TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2}
	//root1.Right = &TreeNode{Val: 3}
	root2 := TreeNode{Val: 1}
	//root2.Left = &TreeNode{Val: 2}
	root2.Right = &TreeNode{Val: 3}

	fmt.Println(isSameTree(&root1, &root2))
}
