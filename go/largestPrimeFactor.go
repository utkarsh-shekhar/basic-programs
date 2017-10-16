package main
import "fmt"

func Find(A ,B ,C int) int {
	if A == 1{
		return C
	}
	if A % B == 0 {
		if B  > C {	//If the current Prime is greater than last prime (C) is greater  
			C = B   // Replace the current greatest integer with the current one
			A = A/B // reduce A to half
		}
		B = 2           //reset B to 2 and start again
	} else {
			B += 1 //Increment B if A is not divided
	}
	return Find(A,B,C)
}

func main() {
	A := 566678909//works for even large numbers mind the stack size
	B := 2
	var C int
	fmt.Println(Find(A,B,C))
}
