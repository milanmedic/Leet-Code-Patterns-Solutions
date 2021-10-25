package main

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

func PreOrderTraversal(node *TreeNode, allNodes []int) []int {
	if node != nil {
		allNodes = append(allNodes, node.value)
		if node.leftChild != nil {
			allNodes = PreOrderTraversal(node.leftChild, allNodes)
		}
		if node.rightChild != nil {
			allNodes = PreOrderTraversal(node.rightChild, allNodes)
		}
	}
	return allNodes
}

func LevelOrderTraversal(node *TreeNode, allNodes []*TreeNode) []*TreeNode {
	if len(allNodes) == 0 {
		allNodes = append(allNodes, node)
	}

	for i := 0; i < len(allNodes); i++ {
		if allNodes[i].leftChild != nil {
			allNodes = append(allNodes, allNodes[i].leftChild)
		}
		if allNodes[i].rightChild != nil {
			allNodes = append(allNodes, allNodes[i].rightChild)
		}
	}

	return allNodes
}

func min(root *TreeNode) int {
	for root.Left != nil {
		root = root.Left
	}
	return root.Val
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	} // key not found or empty tree

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		// key found
		if root.Left == nil { // no child or only right child
			return root.Right
		} else if root.Right == nil {
			return root.Left // only left child
		} else { // both children present
			root.Val = min(root.Right)                    // replace node's value with successor
			root.Right = deleteNode(root.Right, root.Val) // remove the duplicate key from right subtree
		}
	}

	return root
}
