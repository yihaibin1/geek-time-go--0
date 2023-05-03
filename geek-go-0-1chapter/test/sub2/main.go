package main

import "fmt"

func main(){
	c:=&Child{
		parent:parent{},
	}
	c.SayHello()
}


type parent struct {

}

func (p *parent)SayHello(){
	fmt.Println("I am ",p.Name())
}

func (p *parent)Name()string{
	return "parent"
}


type Child struct {
	parent
}

func (c *Child)SayHello(){
	fmt.Println("I am child",c.parent.Name())
}

func (c *Child)Name()string{
	return "hahaha"
}
