package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewNode(value int) *ListNode {
	return &ListNode{Val: value}
}

func (l *ListNode) AddNode(newNode *ListNode) {
	curr := l

	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = newNode
}

func (l *ListNode) CheckForCycle() bool {
	var visitedNodes map[int]bool = make(map[int]bool)
	return hasCycle(l, visitedNodes)
}

func hasCycle(head *ListNode, visitedNodes map[int]bool) bool {
	if head == nil {
		return false
	}
	if _, found := visitedNodes[head.Val]; !found {
		visitedNodes[head.Val] = true
		return hasCycle(head.Next, visitedNodes)
	} else {
		return true
	}
}

func middleNode(head *ListNode) *ListNode {
	len := GetLength(head) / 2

	for len > 0 {
		len--
		head = head.Next
	}

	return head
}

func GetLength(head *ListNode) int {
	len := 0
	for head != nil {
		len++
		head = head.Next
	}

	return len
}

func main() {
	head := &ListNode{Val: 1}
	head.AddNode(NewNode(2))
	head.AddNode(NewNode(3))
	head.AddNode(NewNode(4))
	tails := NewNode(5)
	head.AddNode(tails)

	fmt.Println(middleNode(head))
}
