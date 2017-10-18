package main

import "fmt"
import "time"

func main(){
	i:=2
	fmt.Println("write", i, " as ")
	switch i {/* Basic switch*/
	
	case 1:
		fmt.Println("one")
	
	case 2:
		fmt.Println("two")
	
	case 3:
		fmt.Println("three")
	}
	
	switch time.Now().Weekday(){
		
	case time.Saturday, time.Sunday: 	/* Multiple expressions in same case statement*/
		fmt.Println("it's a weekend")
	default:
		fmt.Println("it's weekday")
	}

	t:=time.Now()
	switch {   	 /* switch without an expression is an alternate way to express if/else logic*/
	
	case t.Hour() <12: 	/* case statements can be non constants also*/
		fmt.Println("it's before noon")	
	
	default:
		fmt.Println("it's after noon")	
	}

}
