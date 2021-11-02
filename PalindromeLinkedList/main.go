package main

import (
	"fmt"
	"math"
)

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

func isPalindrome(head *ListNode) bool {
	ll := ""
	curr := head

	for curr != nil {
		ll = ll + fmt.Sprint(curr.Val)
		curr = curr.Next
	}

	llen := int(math.RoundToEven(float64(len(ll) / 2)))

	for i, j := 0, len(ll)-1; i <= llen && j >= llen; i, j = i+1, j-1 {
		if ll[i] != ll[j] {
			return false
		}
	}

	return true
}

func main() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 2}
	n5 := &ListNode{Val: 1}

	n1.AddNewNode(n2)
	n1.AddNewNode(n3)
	n1.AddNewNode(n4)
	n1.AddNewNode(n5)

	isPalindrome(n1)
}
