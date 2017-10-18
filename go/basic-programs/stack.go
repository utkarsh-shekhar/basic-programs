package main

import "fmt"


type stackdata struct {
	val interface{} // similar void * in C
	next *stackdata // sothat we can keep tarck of next

}

type stack struct {
	top *stackdata //keep track of top of the stack
}

func (k *stack) Push(data interface{}) {
	var l stackdata
	fmt.Println("Inserting at top")
	l.val = data
	l.next = k.top 
	k.top = &l
}

func (k *stack) Pop() interface {} {
	if k.top== nil {
		return "Empty stack"
	}
	fmt.Println("Returning from top")
	v := k.top.val
	k.top = k.top.next
	return v

}
//Implementation Hiding
type Stack interface {
	Push(interface{})
	Pop() interface{}
}
func NewStack() Stack {
	return &stack{}
}

func main() {
	var z = NewStack()
	z.Push(20)
	z.Push(50)
	z.Push(40)
	fmt.Println(z.Pop())
	

}
