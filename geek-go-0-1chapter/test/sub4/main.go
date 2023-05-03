package main

import (
	"fmt"
	"sync"
)

var one sync.Once


func PrintHello(){
	one.Do(func(){
		fmt.Println("Hello")
	})
}

func main(){
	PrintHello()
}