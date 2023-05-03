package server

import "net/http"

type Routeable interface {
	Route(method string,pattern string, handleFunc handlerFunc)
}

type Handler interface {
	ServeHTTP(c *Context)
	Routeable
}

type HandlerBasedMap struct {
	handlers map[string]func(ctx *Context)
}

func (h *HandlerBasedMap)Route(method string,
	pattern string,
	handleFunc handlerFunc) {
	key:=h.key(method,pattern)
	h.handlers[key]=handleFunc
}

func (h *HandlerBasedMap) ServeHTTP(c *Context) {
	key:=h.key(c.R.Method,c.R.URL.Path)
	if handler,ok:=h.handlers[key];ok{
		handler(c)
	}else{
		c.W.WriteHeader(http.StatusNotFound)
		c.W.Write([]byte("NOT FOUND"))
	}
}

func (h *HandlerBasedMap)key(method string,pattern string)string{
	return method+"#"+pattern
}


//确保接口的实现，无特殊含义
var _ Handler=&HandlerBasedMap{}

func NewHandlerBasedMap()Handler{
	return &HandlerBasedMap{
		handlers: make(map[string]func(ctx *Context)),
	}
}
