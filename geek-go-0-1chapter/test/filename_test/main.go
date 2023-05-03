package main

import (
	"fmt"
	"os"
)

func main(){
	f,_:=os.Open("main.go")
	b:=make([]byte,100000)
	f.Read(b)
	fmt.Println(string(b))
}
