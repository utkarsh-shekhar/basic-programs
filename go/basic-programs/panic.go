package main

import "fmt"

func badfunction() {
	fmt.Printf("Select Panic type (0=no panic, 1=int, 2=runtime panic)\n")
	var choice int
	fmt.Scanf("%d", &choice)
	switch choice {
		case 1:
			panic(0)
		case 2:
			var invalid func()
			invalid()
	}	
}

func main() {
	defer func() {
		if x:= recover(); x!= nil {
			switch x.(type){
				default:
					 panic(x)
				case int:
					fmt.Printf("function panicked with a very helpful error: %d\n", x)
			}
		}
	}()
	badfunction()
	fmt.Printf("Program exited normally\n")
}

