package main

import (
	"fmt"
	"unicode"
	"strings"
	"golang.org/x/exp/utf8string"
	)

func main(){
	str:= "\t the important rôle of utf8 text\n"
	str = strings.TrimFunc(str, unicode.IsSpace)
	
	//The wrong way	
	fmt.Printf("%s\n", str[0: len(str)/2])
	//The right way
	u8 := utf8string.NewString(str)
	FirstHalf := u8.SLice(0, u8.RuneCount()/2)
	fmt.Printf("%s\n",FirstHalf)

	
}

