// Package parser 实现栈结构
package parser

import (
	"container/list"
)

// Stack represents a LIFO (Last In First Out) data structure
// 使用Go标准库中的list实现栈结构
type Stack struct {
	list *list.List
}

// NewStack creates and initializes a new Stack
// 创建并初始化一个新的栈
func NewStack() *Stack {
	return &Stack{list: list.New()}
}

// pop removes and returns the top element from the stack
// 从栈顶移除并返回元素，如果栈为空则返回nil
func (stack *Stack) pop() interface{} {
	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

// push adds a new element to the top of the stack
// 将新元素添加到栈顶
func (stack *Stack) push(v interface{}) {
	stack.list.PushBack(v)
}

// isEmpty checks if the stack has no elements
// 检查栈是否为空
func (stack *Stack) isEmpty() bool {
	return stack.list.Len() == 0
}

// top returns the top element without removing it from the stack
// 返回栈顶元素但不移除它，如果栈为空则返回nil
func (stack *Stack) top() interface{} {
	e := stack.list.Back()
	if e != nil {
		return e.Value
	}
	return nil
}
