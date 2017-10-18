package main

import (
	"fmt"
	)

type tree struct{
	left,right *tree
	data int
}

//inorder traversal
func traverse(t *tree){   // passing the root and channel
	fmt.Println("in traverse",t.data)
	if t == nil {
		return
	}else {
		traverse(t.left)
		fmt.Println(t.data)
		traverse(t.right)
	}
}

func insert(t *tree, l *tree) *tree {
	if t == nil{
		return t
	}
	if l.data < t.data {
		fmt.Println("Inserting into left")
		insert(t.left,l)
	}else{
		fmt.Println("Inserting into right")
		insert(t.right,l)
	}
	return t
	
}
	


func main() {
	var  t = new(tree)
	k := 80
	t.data = k
	var  newnode = new(tree)
	newnode.data = 20
	insert(t,newnode)
	newnode.data = 90
	insert(t,newnode)
	traverse(t)

}

