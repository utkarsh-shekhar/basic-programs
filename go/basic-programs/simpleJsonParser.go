package main

import ("fmt"
		"strings"
		//"os"
		//"bufio"
		"regexp"
		"io/ioutil"
		
		)
		

func jsonObjectParser(jsonData string)(string , []string) {
	fmt.Println("inside jsonObjectParser")

	parseString := make([]string,0)
	jsonData = jsonData[1:]
	var result string
	jsonData = strings.TrimSpace(jsonData)
	if jsonData[0] != '}' {
		result,jsonData = stringParser(jsonData)
		parseString = append(parseString,result)
		result,jsonData = colonParser(jsonData)
		
		parseString = append(parseString,result) 
		fmt.Println("appended")
		for jsonData[0] != '}'{
				result,jsonData = elementParser(jsonData)
				parseString = append(parseString, result)
				jsonData = strings.TrimSpace(jsonData)
		}
	}
	fmt.Println(parseString)
	return result,parseString

}
func stringParser(jsonData string) (string,string){
	fmt.Println("Inside String parser")
	var result string
	var index int
	jsonData = strings.TrimSpace(jsonData)
	fmt.Println(jsonData)
	if jsonData[0] =='"'{
		fmt.Println("string found ")
		jsonData = jsonData[1:]
		for jsonData[index] != '"'{
			index++
		}
		fmt.Println(index)
		result = jsonData[:index]
		return result,jsonData[index +1:]
	
		fmt.Println("outside stringParser")
		return result,jsonData
	}	
	return result,jsonData
}

func booleanParser(jsonData string) (string, string) {
	fmt.Println("Inside boolean parser")
	var result string
	jsonData = strings.TrimSpace(jsonData)
	if len(jsonData) > 4{
		if jsonData[0:4] == "true"{
			//fmt.Println("bool found")
			result = "true"
			return result,jsonData[4:]
		}else if jsonData[0:5] == "false" {
			result = "false"
			//fmt.Println("bool found")
			return result,jsonData[5:]
		}
	}
	return result,jsonData
		
}


func colonParser(jsonData string) (string,string) {
	fmt.Println("Inside colon parser")
	var result string
	jsonData = strings.TrimSpace(jsonData)
	result = string(jsonData[0])
	if jsonData[0] == ':'{
		 fmt.Println("colon found")
		 //fmt.Println(jsonData[1:])
		return result,jsonData[1:]
	}
	return result,jsonData
}

func commaParser(jsonData string) (string, string){
	var result string
	fmt.Println("inside comma parser")
	jsonData = strings.TrimSpace(jsonData)
	if jsonData[0] ==','{
		result = ","
		fmt.Println("Comma found")
		fmt.Println(jsonData[1:])
		return result,jsonData[1:]
	}
	return result,jsonData

}

func nullParser(jsonData string) (string, string) {
	fmt.Println("Inside null parser")
	var result string
	jsonData = strings.TrimSpace(jsonData)
	if len(jsonData) > 4{
		if jsonData[0:4] == "null"{
			fmt.Println("null found")
			result = "nil"
			return result,jsonData[4:]
		}
	}
	return result,jsonData
}

func numberParser(jsonData string) ([]string, string) {
	jsonData = strings.TrimSpace(jsonData)
	fmt.Println("inside number parser")
	var result [][]int
	//fmt.Println(result1)
	re:= regexp.MustCompile("^[-+]?[0-9]*.?[0-9]+([eE][-+]?[0-9]+)?")
	result1 := re.FindAllString(jsonData, -1)
	if len(result1) > 0{
		//fmt.Println("inside num if ")
		result = re.FindAllStringIndex(jsonData, -1)
		fmt.Println(result)
		index := result[0][1]
		fmt.Println(index)
		//fmt.Println("Flag 1")
		return result1,jsonData[index:]
	}else{
		//fmt.Println("Flag 2")
		return result1, jsonData
	}
	
}

func jsonArrayParser(jsonData string) (string ,[]string) {
	fmt.Println("inside jsonArrayParser")
	parseString := make([]string,0)
	jsonData = jsonData[1:]
	var result string
	jsonData = strings.TrimSpace(jsonData)
	if jsonData[0] != ']' {
		result,jsonData = stringParser(jsonData)
		parseString = append(parseString,result)
		
		result,jsonData = colonParser(jsonData)
		
		parseString = append(parseString,result) 
		fmt.Println("appended")
		for jsonData[0] != ']'{

				result,jsonData = elementParser(jsonData)
				parseString = append(parseString, result)
				fmt.Println("appended")
				jsonData = strings.TrimSpace(jsonData)
				fmt.Println(jsonData)
			
		}
	
	}
	fmt.Println(parseString)

	return result,parseString
}
func elementParser(jsonData string) (string, string) {
	var result string
	
	fmt.Println("Inside Element Parser")
	jsonData = strings.TrimSpace(jsonData)
	
	result,jsonData = stringParser(jsonData)
	if result != "" {
		return result, jsonData	
	}
	//Problems in number parser
	resultnum,jsonData := numberParser(jsonData)
	if len(resultnum) > 0{
		//fmt.Println(resultnum)
		return resultnum[0],jsonData
	}
	resultbool,jsonData := booleanParser(jsonData)
	if resultbool!= "" {
		if resultbool == "true"{
			return "true",jsonData
			}else {
				return "false",jsonData
		}
		
	}
	resultnull,jsonData := nullParser(jsonData)
	if resultnull != "" {
		return resultnull,jsonData
	}
	resultcomm,jsonData := commaParser(jsonData)
	if resultcomm != "" {
		fmt.Println(jsonData)
		return resultcomm,jsonData
	}/*
	resultarr,jsonData :=jsonArrayParser(jsonData)
	if resultarr != ""{
		return resultarr,jsonData
	}
	resultobj,jsonData :=jsonObjectParser(jsonData)
	if resultobj != ""{
		return resultobj,jsonData
	}*/
	
	resultcol,jsonData := colonParser(jsonData)
	if resultcol != "" {
		return resultcol,jsonData
	}
	return result,jsonData
}

func main(){
	//file, err := os.Open(s)
	
	buf,err := ioutil.ReadFile("test.txt")// whole file at a time (Readall) 
	if err != nil {
		fmt.Println(err)
	}
	data:= string(buf)
	fmt.Println(data)
	if data[0] == '{' {
		strings.TrimSpace(data)
		fmt.Println(data)
				jsonObjectParser(data[:])
	}else if data[0] =='['	{
		strings.TrimSpace(data)
		jsonArrayParser(data[:])
	}
	
	

}
