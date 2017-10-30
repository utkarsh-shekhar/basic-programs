package main

import "fmt"
import "os"


func main(){
	file,err := os.Open("switch.go")
	if err != nil {
		fmt.Println(err)
	}
	data := make([]byte, 100)
	for n, err := file.Read(data); err == nil; n,err = file.Read(data){ // reading into buff until EOF
		if n > 0{
			os.Stdout.Write(data)
		}
	}
}