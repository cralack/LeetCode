package flattennestedlistiterator

import (
	"testing"
)

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
	Val  int
	List []*NestedInteger
	It   NestedIterator
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool {
	return this.List == nil
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {
	return this.Val
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {

}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {

}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {
	return this.List
}

type NestedIterator struct {
	Stack []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	stack := []*NestedInteger{}
	for i := len(nestedList) - 1; i >= 0; i-- {
		stack = append(stack, nestedList[i])
	}
	return &NestedIterator{Stack: stack}
}

func (this *NestedIterator) Next() int {
	this.stackTop2Int()
	top := this.Stack[len(this.Stack)-1]
	this.Stack = this.Stack[:len(this.Stack)-1]
	return top.GetInteger()
}

func (this *NestedIterator) HasNext() bool {
	this.stackTop2Int()
	return len(this.Stack) > 0
}
func (this *NestedIterator) stackTop2Int() {
	for len(this.Stack) > 0 {
		top := this.Stack[len(this.Stack)-1]
		if top.IsInteger() {
			return
		}
		this.Stack = this.Stack[:len(this.Stack)-1]
		list := top.GetList()
		for i := len(list) - 1; i >= 0; i-- {
			this.Stack = append(this.Stack, list[i])
		}
	}
}
func Test_flatten_nested_list_iterator(t *testing.T) {
}
