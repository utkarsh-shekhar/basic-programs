package main 

import "fmt"


func Reduce(data []int,f func (int, int) int,curr int, index int,result int) interface	{}{
	if index > (len(data) -1){
		fmt.Println("Recursion done :)")
		return result
	}
	
	curr =  index
	//fmt.Println(curr)
	result = f(data[curr],result)
	//fmt.Println(result)
	index++
	return Reduce(data[:],f,curr,index,result
}

func main() {
	arr:= []int{1,2,3,4,5,6,7,8}
	fmt.Println(arr)	

	var result int
	function:= func (a int,b int) int{
		return a+b
	}
	fmt.Println(Reduce(arr[:],function,0,0,result))


}