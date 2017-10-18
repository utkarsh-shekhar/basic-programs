package main

import "fmt"

type node struct {
	left *node
	right *node
	data  int
}


func newnode(val int) *node{
	var newnode = new(node)
	newnode.left = nil
	newnode.right = nil
	newnode.data = val
	return newnode
}

func main(){
	var root= new(node) 
	root = newnode(10)
	root.left = newnode(6)
	root.right = newnode(12)
	root.left.right = newnode(7)
	root.right.left = newnode(11)
	fmt.Println(root)
	fmt.Println(root.left)
	fmt.Println(root.right)

	fmt.Println(root.left.right)
	fmt.Println(root.right.left)


}