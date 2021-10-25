package main

import (
	"fmt"
	"math"
)

type BinaryTree struct {
	root *TreeNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (tree *BinaryTree) Add(value int) {
	tree.root = Traverse(tree.root, value)
}

func Traverse(node *TreeNode, value int) *TreeNode {
	if node == nil {
		node = &TreeNode{Val: value, Left: nil, Right: nil}
		return node
	}

	if node.Val < value {
		node.Right = Traverse(node.Right, value)
		return node
	} else {
		node.Left = Traverse(node.Left, value)
		return node
	}
}

func main() {
	var bt BinaryTree = BinaryTree{root: nil}
	bt.Add(12)
	bt.Add(5)
	bt.Add(2)
	bt.Add(9)
	bt.Add(18)
	bt.Add(14)
	bt.Add(19)
	bt.Add(13)
	bt.Add(17)
	bt.Add(16)
	bt.Add(15)
	fmt.Println(isSymmetric(bt.root))
}

func isSymmetric(root *TreeNode) bool {
	allNodes := &[][]TreeNode{}
	LvlOrderTraverse(root, allNodes, 1)

	for i := 0; i < len(*allNodes); i++ {
		for j := 0; j < len((*allNodes)[i]); j++ {
			k := len((*allNodes)[i]) - j - 1
			if (*allNodes)[i][j].Val != (*allNodes)[i][k].Val {
				return false
			}
		}
	}

	return true
}

func LvlOrderTraverse(node *TreeNode, allNodes *[][]TreeNode, level int) {
	if node != nil {
		if len(*allNodes) < level {
			*allNodes = append(*allNodes, []TreeNode{})
		}

		(*allNodes)[level-1] = append((*allNodes)[level-1], *node)
		level++
		LvlOrderTraverse(node.Left, allNodes, level)
		LvlOrderTraverse(node.Right, allNodes, level)
	} else {
		if len(*allNodes) < level {
			*allNodes = append(*allNodes, []TreeNode{})
		}

		(*allNodes)[level-1] = append((*allNodes)[level-1], TreeNode{Val: math.MaxUint32})
	}
}
