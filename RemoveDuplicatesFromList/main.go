package main

const MAX_NUM = 9999999999999999

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

func deleteDuplicates(head *ListNode) *ListNode {
	curr := head

	for curr != nil {
		if (curr.Next != nil) && curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}

func main() {
	n1 := &ListNode{Val: 1}
	n5 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n4 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}

	n1.AddNewNode(n5)
	n1.AddNewNode(n2)
	n1.AddNewNode(n4)
	n1.AddNewNode(n3)

	deleteDuplicates(n1)
	print("Endee")
}
