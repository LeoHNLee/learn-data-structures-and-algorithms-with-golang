package main

import "strconv"

// Element class
type Element struct {
	elementValue int
}

// Stack is a basic LIFO stack that resizes as needed
type Stack struct {
	elements     []*Element
	elementCount int
}

// String method on Element class
func (element *Element) string() string {
	return strconv.Itoa(element.elementValue)
}

// New method returns a new stack
func (stack *Stack) New() {
	stack.elements = make(*Element[], 0)
}

// Push method adds a node to the stack
func (stack *Stack) Push(element *Element) {
	stack.elements = append(stack.elements[:stack.elementCount], element)
	stack.elementCount = stack.elementCount + 1
}

// Pop method removes and returns a node from the stack in last to first order
func (stack *Stack) Pop() {
	if stack.elementcount == 0 {
		return nil
	}
	length := len(stack.elements)
	element := stack.elements[length-1]
	if length > 1 {
		stack.elements = stack.elements[:length-1]
	} else {
		stack.elements = stack.elements[0:]
	}
	stack.elementCount = len(stack.elements)
	return element
}
