package main

import (
	"fmt"
	"net/http"
	"project1/myWebFrame/server"
)


func main() {
	myServer := server.NewServer("Yihaibin")
	myServer.Route(http.MethodPost,"/sign/up", server.Sign)
	defer func(){
		if data:=recover();data!=nil{
			fmt.Println("恢复错误:",data)
		}
		myServer.Start(":8080")
	}()
	panic(http.ErrServerClosed)
	myServer.Start(":8080")
}


