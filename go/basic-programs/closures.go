package main 

import "fmt"


func closure() func () int{
	a:= 10
	return func() int{
		return a*a
	}
	
}

func main(){
	k:= closure()
	fmt.Println(k())
	
}