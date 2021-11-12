package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	min := 9999999
	pmin := &min
	Traverse(root, pmin, 0)
	return min
}

func Traverse(node *TreeNode, min *int, currMin int) {
	if node != nil {
		currMin++
		if node.Left != nil {
			Traverse(node.Left, min, currMin)
		}

		if node.Right != nil {
			Traverse(node.Right, min, currMin)
		}

		if node.Left == nil && node.Right == nil {
			if currMin < *min {
				*min = currMin
			}
		}

		return
	}
	if currMin < *min {
		*min = currMin
	}
}

//[2,1,3,7,4,9,5,8,6]
func main() {
	root := TreeNode{Val: 2}
	n1 := TreeNode{Val: 1}
	root.Left = &n1
	n2 := TreeNode{Val: 3}
	root.Right = &n2
	n3 := TreeNode{Val: 7}
	n1.Left = &n3
	n4 := TreeNode{Val: 4}
	n1.Right = &n4
	n7 := TreeNode{Val: 8}
	n3.Left = &n7
	n8 := TreeNode{Val: 6}
	n3.Right = &n8
	n5 := TreeNode{Val: 9}
	n2.Left = &n5
	n6 := TreeNode{Val: 5}
	n2.Right = &n6

	minDepth(&root)
}
