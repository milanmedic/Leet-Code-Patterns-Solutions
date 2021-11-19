package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return Traverse(root, targetSum, 0)
}

func Traverse(node *TreeNode, targetSum int, currSum int) bool {
	if node != nil {
		currSum += node.Val

		if node.Left == nil && node.Right == nil {
			if targetSum == currSum {
				return true
			} else {
				return false
			}
		}

		return (Traverse(node.Left, targetSum, currSum) || Traverse(node.Right, targetSum, currSum))
	}
	if currSum == targetSum {
		return true
	}
	return false
}

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 8}

	root.Left.Left = &TreeNode{Val: 11}
	root.Right.Left = &TreeNode{Val: 13}
	root.Right.Right = &TreeNode{Val: 4}

	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Left.Left.Right = &TreeNode{Val: 2}

	root.Right.Right.Right = &TreeNode{Val: 1}

	hasPathSum(root, 26)
}
