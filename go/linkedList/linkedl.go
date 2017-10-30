package  main

	
import "fmt"

type link struct{
	data int
	next *link
}

func printlist(root *link){
	var temp *link
	temp =  root
	if(root == nil){
		return
	}else{
		for temp.next != nil{
			fmt.Println(temp.data)
			temp = temp.next
		}
	}
}

func insertlistH(root *link, value int) *link{

	var newnode  *link
	newnode.data = value
	newnode.next = root
	root = newnode
	return root
}

func main(){
	var t link
	t.data = 5
	t.next = nil

	insertlistH(&t,9)
	printlist(&t)
	insertlistH(&t,10)
	printlist(&t)
	//fmt.Println("Linked list contents: ", t.data)
}