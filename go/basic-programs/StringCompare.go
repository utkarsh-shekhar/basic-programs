package main

import "fmt"
import "strings"

func main(){
	str1:= "A string"
	str2:= "A" + "string"

	if str1 == str2{
		fmt.Println("str1 and str2 are equal")
	}
	if (&str1 == &str2) {
		fmt.Println("str1 and str2 are identical")
	}
	str2 += "with a suffix"
	if (str1 < str2){
		fmt.Println("str1 comes before str2")
	}
	str := "Etoile"
	// DO this for iterating over strings
	for _,c := range str {
		fmt.Printf("%c",c)
	}
	str3 := "\t This is a string \n"
	str3 = strings.Trim(str3, "\t\n\r")
	fmt.Printf("%s\n",str3)
	words := strings.Split(str3, " ")
	for _, word := range words {
		fmt.Printf("%s\n",word)
	}
}
