package server

import (
	"errors"
	"net/http"
	"strings"
)

var ErrorPatter=errors.New("invalid route way")

func NewHandlerBasedOnTree()Handler{
	return &HandlerBasedOnTree{
		root:&node{},
	}
}


type HandlerBasedOnTree struct {
	root *node
}

type node struct {
	//按照视频的意思，存储的不是完整的路径
	path string
	children []*node
	//如果这是叶子结点
	//那么匹配上后就可以调用该方法
	handler handlerFunc

}

func (h *HandlerBasedOnTree) ServeHTTP(c *Context) {

	handler,found:=h.findRouter(c.R.URL.Path)
	if found{
		handler(c)
	}else{
		c.W.WriteHeader(http.StatusNotFound)
		_, _ = c.W.Write([]byte("Not Found"))
		return
	}

}

func(h *HandlerBasedOnTree)findRouter(path string)(handlerFunc,bool){
	paths:=strings.Split(strings.Trim(path,"/"),"/")
	cur:=h.root
	for _,nowPath:=range paths{
		node,found:=h.findMatchChild(cur,nowPath)
		if !found{
			return nil,false
		}
		cur=node
	}
	if cur.handler==nil{
		return nil,false
	}
	return cur.handler,true
}

func (h *HandlerBasedOnTree) Route(method string,
	pattern string,
	handleFunc handlerFunc)error{

	err:=h.validPatternTest(pattern)
	if err!=nil{
		return err
	}
	pattern=strings.Trim(pattern,"/")
	paths:=strings.Split(pattern,"/" )
	cur:=h.root
	for index,path:=range paths{
		mathChild,ok:=h.findMatchChild(cur,path)
		if ok{
			cur=mathChild
		}else{
			h.createSubTree(cur,paths[index:],handleFunc)
		}
	}
	cur.handler=handleFunc
	return nil
}

func (h *HandlerBasedOnTree)findMatchChild(cur *node,
	path string)(*node,bool){
	var wildCardNode *node
	for _,child:=range cur.children{
		if child.path==path&&child.path!="*"{
			return child,true
		}
		if child.path=="*"{
			wildCardNode=child
		}
	}
	return wildCardNode,wildCardNode!=nil

}

func(h *HandlerBasedOnTree)validPatternTest(pattern string)error{
	index:=strings.Index(pattern,"*")
	if index>=0{
		if index==0{
			return ErrorPatter
		}
		if index<len(pattern)-1{
			return ErrorPatter
		}
		if pattern[index-1]!='/'{
			return ErrorPatter
		}
	}
	return nil
}



func (h *HandlerBasedOnTree)createSubTree(root *node,
	paths []string,
	handleFunc handlerFunc){
	cur:=root
	for _,path:=range paths{
		nn:=h.NewNode(path)
		cur.children=append(cur.children,nn)
		cur=nn
	}
	cur.handler=handleFunc
}

func(h *HandlerBasedOnTree)NewNode(path string)*node{
	return &node{
		path:     path,
		children: make([]*node,0,2),
	}
}



/*func (n *HandlerBasedOnTree) Route(method string,
	pattern string,
	handleFunc func(ctx *Context)) {
	pattern=strings.Trim(pattern,"/")
	elements:=strings.Split(pattern,"/" )
	find(elements,n.root,0,handleFunc)

}

// find 插入最新叶子结点
func find(elements []string,nowNode *node,level int,handler handlerFunc){
	if len(elements)==0{
		return
	}
	if len(elements)==level{
		nowNode.handler=handler
		return
	}
	for _,val:=range nowNode.children{
		path:=strings.Trim(val.path,"/")
		paths:=strings.Split(path,"/")
		if paths[level]==elements[level]{
			find(elements,val,level+1,handler)
		}
	}
	var newPath string
	for i:=0;i<=level;i++{
		newPath=fmt.Sprintf(newPath+"/%s",elements[i])
	}
	createNode:=&node{
		path: newPath,
		children:[]*node{},
		handler:nil,
	}
	nowNode.children=append(nowNode.children,createNode)
	find(elements,createNode,level+1,handler)
}*/