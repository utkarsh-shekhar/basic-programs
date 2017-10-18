package main

import "fmt"

func main(){
	if 7%2 == 0{
		fmt.Println("7 is Even")
	}else{
		fmt.Println("7 is Odd")
	}
	
	if num:= 9/*declare and initialize statement*/; num < 0/* conditional can be preceded by by statements*/ {
		fmt.Println(num, "is negative")
	}else if num < 10{
		fmt.Println(num, "has 1 digit")
	}else {
		fmt.Println(num, "has multiple digits")
	}
	/* There is no Ternary if else in go i.e. ?:*/
}	
