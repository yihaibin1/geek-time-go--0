package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}


func (c *Context)ReadJson(sreq interface{})error{
	req:=c.R
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, sreq)
	fmt.Println(sreq)
	if err != nil {
		return err
	}
	return nil
}

func (c *Context)WriteJson(code int, sres interface{}) error {
	c.W.WriteHeader(code)
	respResult,err:=json.Marshal(sres)
	if err!=nil{
		return err
	}
	_,err=c.W.Write(respResult)
	return err
}

func (c *Context)OkJson(sres interface{})error{
	return c.WriteJson(http.StatusOK,sres)
}



func NewContext(W http.ResponseWriter,R *http.Request)*Context{
	return &Context{
		W: W,
		R: R,
	}
}



/*func NewContext(handleFunc func(ctx *Context))func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		ctx:=&Context{
			W: w,
			R: r,
		}
		handleFunc(ctx)
	}
}*/