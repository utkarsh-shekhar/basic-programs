package main

import( "fmt"
	"encoding/json"
	)
/*
type Message struct{
	Name string	//Declared with caps letter for Public access
	Body string
	Time int64
}


type HalfMessage struct {
	Name string
	Body string
}  
*/
func jsonDecoder(data[] byte) {

	switch data {
	case data[0] == '{':
		objJson(data)
	case data[0] =='[':
		arrayJson(data)
	}

}

func BoolParse(data) {
	
}

func Delimitors(data){

}
func NumParse(data) {
	
}

func StringParse(data) {
	 var obj string
	 
}
func arrayJson(data) {
	
}
func objJson(data) {
	obj	
}
func NullParse(data){

}

func main(){	
	m:= Message{"Alice", "Hello", 12345 }
	b,err:= json.Marshal(m)
	if err == nil {
		fmt.Println("json encoding successfull")
	}
	/*var k Message
	fmt.Printf("%s \n", b)
	
	err = json.Unmarshal(b, &k)
	fmt.Println(k)*/
	
	//Unmarshal the json data in different structure
	/*
	var h HalfMessage
	json.Unmarshal(b,&h)
	fmt.Println(h)
	*/
	//if we don't know the data structure
	//Unmarshal the json data in an interface{} which returns  map["key" as string]interface{} as"value"
	/*
	var f interface{}
	json.Unmarshal(b,&f)
	fmt.Println(f)
	*/



}


