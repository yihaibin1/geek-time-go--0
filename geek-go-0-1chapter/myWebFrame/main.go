package main

import (
	"fmt"
	"log"
	"net/http"
	"project1/myWebFrame/server"
)


func main() {
	myServer := server.NewServer("Yihaibin")
	err:=myServer.Route(http.MethodPost,"/sign/*", server.Sign)
	if err!=nil{
		log.Println("route register err:",err)
	}
	defer func(){
		if data:=recover();data!=nil{
			fmt.Println("恢复错误:",data)
		}
		myServer.Start(":8080")
	}()
	panic(http.ErrServerClosed)
	myServer.Start(":8080")
}


