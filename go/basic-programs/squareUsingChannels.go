package main 

import "fmt"


func gen(nums... int) <-chan int{
	out:= make(chan int)  // making the channel
	go func() { 
		for _,i := range nums{
			out <-i  // writing into channel
		}
		close(out) // closing the channel
	}()
	
	return out
}

func sq(in <-chan int)<- chan int{ // passing arguments from channel into function
	out:= make(chan int)
	go func() {            // go routine
		for n:= range in{
			out<- n *n     // writing into channel
		}
		close(out)
	}()
	
	return out
}

 func main() {
		c:=gen(2,3)
		out := sq(c)
		fmt.Println(<-out)

}