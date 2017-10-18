package main

import "fmt"
type Any interface{}

func main(){
	a:= make(map[int] string)
	b:= make(map[Any] int )

	a[12] = "A string"
	a[12.0] = "B string"
	b[23] = 23
	fmt.Printf("%s %s %d\n", a[12], a[12.0],b[23])
}
