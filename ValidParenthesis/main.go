package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(IsValid("([)]"))
	fmt.Println(IsValid("()[]{}"))
	fmt.Println(IsValid("]"))
}

type Stack struct {
	top      *string
	contents []string
}

func CreateStack() Stack {
	var stack Stack = Stack{}
	stack.contents = []string{}
	return stack
}

func (stack *Stack) Push(newElem string) string {
	stack.contents = append(stack.contents, newElem)
	stack.top = &stack.contents[len(stack.contents)-1]
	return *stack.top
}

func (stack *Stack) Peek() string {
	if stack.top != nil {
		return *stack.top
	}
	return ""
}

func (stack *Stack) Pop() string {
	var poppedElem = stack.contents[len(stack.contents)-1]
	stack.contents = stack.contents[0 : len(stack.contents)-1]
	if len(stack.contents) > 0 {
		stack.top = &stack.contents[len(stack.contents)-1]
	} else {
		stack.top = nil
	}

	return poppedElem
}

func (stack *Stack) IsEmpty() bool {
	return len(stack.contents) <= 0
}

func IsValid(parenthesis string) bool {
	var stack Stack = CreateStack()
	for i := 0; i < len(parenthesis); i++ {
		if CheckForParenType(parenthesis[i], "{") || CheckForParenType(parenthesis[i], "[") || CheckForParenType((parenthesis[i]), "(") {
			stack.Push(string(parenthesis[i]))
		} else if CheckForParenType(parenthesis[i], "}") || CheckForParenType(parenthesis[i], "]") || CheckForParenType(parenthesis[i], ")") {
			if stack.Peek() == "{" && CheckForParenType(parenthesis[i], "}") {
				if !stack.IsEmpty() {
					stack.Pop()
				} else {
					return false
				}
			} else if stack.Peek() == "[" && CheckForParenType(parenthesis[i], "]") {
				if !stack.IsEmpty() {
					stack.Pop()
				} else {
					return false
				}
			} else if stack.Peek() == "(" && CheckForParenType(parenthesis[i], ")") {
				if !stack.IsEmpty() {
					stack.Pop()
				} else {
					return false
				}
			} else {
				return false
			}
		}
	}

	return stack.IsEmpty()
}

func CheckForParenType(character byte, parenType string) bool {
	return strings.EqualFold(string(character), parenType)
}
