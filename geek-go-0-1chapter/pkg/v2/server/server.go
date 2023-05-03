package server

import "net/http"

type Routeable interface {
	Route(method string,pattern string, handleFunc handlerFunc)error
}

type Handler interface {
	ServeHTTP(c *Context)
	Routeable
}

type Server interface {
	Routeable
	Start(address string) error
}

func NewServer(name string,builders...FilterBuild) Server {
	handler:=NewHandlerBasedOnTree()
	var root Filter= handler.ServeHTTP

	for i:=len(builders)-1;i>=0;i--{
		b:=builders[i]
		root=b(root)
	}
	return &sdkHttpServer{
		Name: name,
		Handler:handler,
		root:root,
	}
}

type sdkHttpServer struct {
	Name string
	Handler Handler
	root Filter
}

func (s *sdkHttpServer) Route(method string,
	pattern string,
	handleFunc handlerFunc)error {
	return s.Handler.Route(method,pattern,handleFunc)
}

func (s *sdkHttpServer) Start(address string) error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		c:=NewContext(writer,request)
		s.root(c)
	})
	return http.ListenAndServe(address, nil)
}
