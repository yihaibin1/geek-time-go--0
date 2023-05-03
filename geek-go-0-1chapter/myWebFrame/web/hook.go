package web

import (
	"context"
	"errors"
	"fmt"
	"project1/myWebFrame/server"
	"sync"
)

type Hook func(ctx context.Context)error


func BuildCloseServerHook(servers ...server.Server)Hook{
	return func(c context.Context)error{
		wg:=sync.WaitGroup{}
		wg.Add(len(servers))

		DoneSigure:=make(chan struct{})

		for _,srv:=range servers{
			go func(srv server.Server){
				err:=srv.Shutdown(c)
				if err!=nil{
					fmt.Println("showdown failed:",err)
				}
				wg.Done()
			}(srv)
		}

		go func(){
			wg.Wait()
			DoneSigure<- struct{}{}
		}()

		select {
			case <-DoneSigure:
				fmt.Println("所有服务正常关闭")
				return nil
			case <-c.Done():
				fmt.Println("关闭超时，已经强制关闭")
				return errors.New("关闭超时，已强行关闭")
		}
	}
}