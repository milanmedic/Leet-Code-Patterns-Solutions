package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) AddNewNode(nn *ListNode) {
	curr := n

	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = nn
}

func reverseList(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	return prev
}

func main() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}

	n1.AddNewNode(n2)
	n1.AddNewNode(n3)
	n1.AddNewNode(n4)
	n1.AddNewNode(n5)

	reverseList(n1)
}
