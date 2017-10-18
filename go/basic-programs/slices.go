package main
import "fmt"

func main(){
	s:= make([]string,3)//slice declaration
	fmt.Println("emp: ",s)
	
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ",s)
	fmt.Println("get: ",s[2])
	
	fmt.Println("len: ",len(s))
	s = append(s,"d","e") //we can append in slices but not in arrays 
	fmt.Println("now: ",s)

	c:=make([]string,len(s))
	copy(c,s)// copy one slice from another
	fmt.Println("cpy: ",c)	

	l:=s[2:5] //just like python 
	m:=s[2:] //just like python 
	n:=s[:5] //just like python  
	fmt.Println("sl1: ", l)
	fmt.Println("sl2: ", m)
	fmt.Println("sl3: ", n)
	


	t := []string{"g","h","i"} // declaration and initialize in same line
	fmt.Println("edf: ",t)		
	
	//two dimensional slices
	twoD :=make([][]int,3)
	for i:=0; i< 3; i++{
		innerlen := i+1	//length of inner slice
		twoD[i]=make([]int,innerlen)//length of inner slices can vary unlike two dimensional arrays
		for j:=0; j< innerlen; j++{
			twoD[i][j]= i+j

		}
	}
	fmt.Println("2d", twoD)
 
}
