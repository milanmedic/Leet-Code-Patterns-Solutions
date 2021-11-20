package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	nodes := []*TreeNode{}
	nodes = FindAppropriateNode(root, subRoot.Val, nodes)
	for _, node := range nodes {
		if TraverseTrees(subRoot, node) {
			return true
		}
	}
	return false
}

func FindAppropriateNode(root *TreeNode, val int, nodes []*TreeNode) []*TreeNode {
	if root != nil {
		if root.Val == val {
			nodes = append(nodes, root)
		}

		nodes = FindAppropriateNode(root.Left, val, nodes)
		nodes = FindAppropriateNode(root.Right, val, nodes)

		return nodes
	}
	return nodes
}

func TraverseTrees(parentNode, childNode *TreeNode) bool {
	if parentNode != nil && childNode != nil {
		if parentNode.Val != childNode.Val {
			return false
		}
		return TraverseTrees(parentNode.Left, childNode.Left) && TraverseTrees(parentNode.Right, childNode.Right)
	} else if parentNode == nil && childNode != nil {
		return false
	} else if parentNode != nil && childNode == nil {
		return false
	}
	return true
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 1}
	//root.Right = &TreeNode{Val: 5}

	//root.Left.Left = &TreeNode{Val: 1}
	//root.Left.Right = &TreeNode{Val: 2}
	//root.Left.Right.Left = &TreeNode{Val: 0}

	subRoot := &TreeNode{Val: 1}
	//subRoot.Left = &TreeNode{Val: 1}
	//subRoot.Right = &TreeNode{Val: 2}

	fmt.Println(isSubtree(root, subRoot))
}
