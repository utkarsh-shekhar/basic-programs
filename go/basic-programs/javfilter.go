package main 

import "fmt"


func Filter(data []int,f func(int) bool, index int, result []interface{}) []interface{}{
	if index > (len(data)-1){
		fmt.Println("Done with recursion")
		return result
	}
	value:= f(data[index])
	if value{
		result = append(result,data[index])
	}
	index++
	return Filter(data[:],f,index,result)
}

func main(){
	even:= func (x int) bool{
		if x%2 == 0 {
			return true
		}
		return false
	}
	arr:= []int{1,2,3,4,5,6,7,8}
	fmt.Println(arr)
	result:= make([]interface{},0)//Slice
	fmt.Println(Filter(arr[:],even,0,result))

}