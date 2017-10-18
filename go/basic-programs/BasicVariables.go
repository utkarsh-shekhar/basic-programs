package main

import "fmt"
import "math"


const s string = "constant string"

func main(){
	fmt.Println("1+1 =",1+1)
	fmt.Println(!true)
	fmt.Println(7.3/2.1)
	var a string = "Init"
	var b,c,d int = 1,2,3						/*Declaring Multiple Variables and initilizing them */
	fmt.Println(a)
	fmt.Println(b, c, d)
	f := "short" 							/* Declaring and initializing variables*/
	fmt.Println(f)
	k:= 1								/* Declaring and initializing variables */ 
	fmt.Println(k)


	const m = 80000000 /* Constant statement can appear any time and a numeric constant has no type until given one such as explicit cast*/
	const n = 3e20/m   /* Constant expressions perform arithematic with arbitary precision*/
	fmt.Println(int64(n))
	
	fmt.Println(n)
	fmt.Println(math.Sin(n)) /*A number can be given a type by using it in a context that rquires one such as variable assignment or function call*/
	
	


	



}
