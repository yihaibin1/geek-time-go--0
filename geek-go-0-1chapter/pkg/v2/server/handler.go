package server

import (
	"log"
	"net/http"
)


//sign 登录请求
func Sign(c *Context) {
	sreq := &signRequest{}
	err:=c.ReadJson(sreq)
	if err!=nil{
		log.Println(err)
	}
	sres := &signResponse{
		UserId: "123",
		Name:   sreq.Name,
	}
	err=c.WriteJson(http.StatusOK,sres)
	if err!=nil{
		log.Println(err)
	}
}