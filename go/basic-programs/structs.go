package main

import "fmt"

type person struct{
	name string
	age int
}

func main(){
	fmt.Println(person{"Bob",20})	//assign value to structure variables omitted fields will be given zero value
	fmt.Println(person{name: "Alice",age: 20})	//assign value to structure variables
	fmt.Println(person{name: "fred"})	//assign value to structure variables
	fmt.Println(&person{name: "Ann",age: 20})	//assign value to structure variables Pointer to the struct

	s:= person{name: "Sean", age: 50}
	fmt.Println(s.name)
	sp:= &s
	fmt.Println(sp.age)// returns a pointer

	sp.age = 51 //modify the contents of sp pointer
	fmt.Println(sp.age)
	
	
}
