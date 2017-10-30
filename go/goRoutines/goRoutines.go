package main

import "fmt"
import "time"
func main(){
	go fmt.Printf("Printed in the background")
	i:= 1
	go fmt.Printf("currently,i is %d\n", i)
	go func(){
		fmt.Printf("i: %d\n",i)
	}()
	i++
	time.Sleep(10000000000)
}
