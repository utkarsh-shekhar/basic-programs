package main

import "fmt"

func main(){
	i:= 1
	for i <= 3{/*basic type*/
		fmt.Println(i)
		i= i+1	
	}

	for j:= 7; j<= 9; j++{/*loop with classical touch*/
		fmt.Println(j)
	}

	for{/* loop infinitly until break*/
		fmt.Println("Loop")
		break
	}
}

