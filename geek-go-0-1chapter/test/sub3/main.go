package main

import (
	"fmt"
	"sync"
)

type student struct {
	information *sync.Map
}

func main(){

	s:=student{information: new(sync.Map)}
	s.information.Store("database",100)
	fmt.Println(s.information.Load("database"))
}
