package main

import (
	"strings"
)

func main(){
	s:="/user/add"
	s=strings.Trim(s,"/")
	elements:=strings.Split(s,"/")
	for _,v:=range elements{
		println(v)
	}
}
