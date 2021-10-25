package main

import "fmt"

func main() {
	var bt Tree = Tree{root: nil}
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
	allNodes := &[][]int{}
	LevelOrderTraversal(bt.root, allNodes, 1)
	fmt.Print(allNodes)
}

type TreeNode struct {
	value      int
	leftChild  *TreeNode
	rightChild *TreeNode
}

type Tree struct {
	root  *TreeNode
	count int
}

func (tree *Tree) Add(value int) {
	tree.root = Traverse(tree.root, value)
	tree.count++
}

func Traverse(node *TreeNode, value int) *TreeNode {
	if node == nil {
		node = &TreeNode{value: value, leftChild: nil, rightChild: nil}
		return node
	}

	if node.value < value {
		node.rightChild = Traverse(node.rightChild, value)
		return node
	} else {
		node.leftChild = Traverse(node.leftChild, value)
		return node
	}
}

func LevelOrderTraversal(node *TreeNode, allNodes *[][]int, level int) {
	if node != nil {
		if len(*allNodes) < level {
			*allNodes = append(*allNodes, []int{})
		}

		(*allNodes)[level-1] = append((*allNodes)[level-1], node.value)
		level++
		LevelOrderTraversal(node.leftChild, allNodes, level)
		LevelOrderTraversal(node.rightChild, allNodes, level)
	}
}
