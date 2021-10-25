package main

import "strings"

//	")()())"

func main() {
	longestValidParentheses("(()")
	longestValidParentheses(")()())")
}

func longestValidParentheses(s string) int {
	var validSubstr string = ""
	var tempValidSubstr string = ""
	var lastOpeningBrace int = 0
	var stack Stack = CreateStack()

	for i := 0; i < len(s); i++ {
		if CheckForParenType(s[i], "(") {
			tempValidSubstr += string(s[i])
			stack.Push(string(s[i]))
			lastOpeningBrace = i
		} else if CheckForParenType(s[i], ")") && stack.Peek() == "(" {
			if !stack.IsEmpty() {
				stack.Pop()
				tempValidSubstr += string(s[i])
			} else {
				validSubstr = tempValidSubstr
			}
		} else {
			tempValidSubstr = ""
			i = lastOpeningBrace
		}
	}

	return len(validSubstr)
}

func CheckForParenType(character byte, parenType string) bool {
	return strings.EqualFold(string(character), parenType)
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
