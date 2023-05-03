package main

import "fmt"

func main(){
	ch:=make(chan int,1)
	close(ch)
	num:=1
	num=<-ch
	fmt.Println(num)

}
