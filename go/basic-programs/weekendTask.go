package main


import "fmt"


func Map(data []int,f func(int) int, index int, result []interface{}) []interface{}{
	if index > (len(data) - 1){
		fmt.Println("Way beyond your limits")
		return result
	}
	//fmt.Println(data[index])
	k:= f(data[index])
	result = append(result,k)
	//fmt.Println(index)
	index++
	return Map(data[:],f,index,result)
	
}


func main(){
	arr:= []int{1,2,3,4,5}
	fmt.Println(arr)
	Square := func (x int) int{
		return x*x
	}
	result := make([]interface{},0)// slice
	//fmt.Println(Square(arr[1]))
	fmt.Println(Map(arr[:],Square,0, result[:]))

}