package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	queue := []TreeNode{*root}
	avgs := []float64{float64(root.Val)}

	Traverse()
}

func Traverse(queue []TreeNode, avgs []float64) []float64 {

}

func main() {

}
