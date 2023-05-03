package web

//优雅关闭所有Http Server
//1.拒绝新的请求：需要一个开关
//2.等待当前所有请求处理完毕：维持请求计数
//3.释放资源：用户释放自己的一些资源
//4.关闭服务器：把所有启动的server都关闭
//5.如果超时，就强制关闭

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"project1/myWebFrame/server"
	"sync"
	"sync/atomic"
	"time"
)

var ErrorHookTimeOut = errors.New("the hook timeout")

type GracefulShutDown struct {
	//还在处理中的请求
	reqCnt int64
	//大于1就说明是要关闭了
	closing int32
	//用channel来通知已经处理完了所有请求
	zeroReqCnt chan struct{}

}

//一旦关闭服务，拒绝接收新请求并处理完老请求
func (g *GracefulShutDown)ShutdownFilterBuilder(next server.Filter)server.Filter{
	return func(c *server.Context){
		cl:=atomic.LoadInt32(&g.closing)
		if cl>0{
			c.W.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		atomic.AddInt64(&g.reqCnt,1)
		next(c)
		n:=atomic.AddInt64(&g.reqCnt,-1)

		if n==0&&cl>0{
			g.zeroReqCnt<- struct{}{}
		}
	}
}


// RejectNewRequestAndWaiting 将会拒绝新的请求，并且等待处理中的请求
func (g *GracefulShutDown)RejectNewRequestAndWaiting(ctx context.Context)error{
	atomic.AddInt32(&g.closing,1)

	if g.reqCnt==0{
		return nil
	}

	done:=ctx.Done()

	select {
		case <-done:
			fmt.Println("处理超时，无法等到所有请求执行完毕")
			return ErrorHookTimeOut
		case <-g.zeroReqCnt:
			fmt.Println("所有请求处理完毕")
	}
	return nil
}


func WaitForShutDown(hooks... Hook){
	signals:=make(chan os.Signal,1)
	signal.Notify(signals,ShutdownSignals...)
	select {
		case <-signals:
			fmt.Println("启动WaitForShutDown，关闭所有服务")

			time.AfterFunc(10*time.Minute,func(){
				fmt.Println("关闭出错，现在进入强制关闭流程")
				os.Exit(1)
			})

			wg:=sync.WaitGroup{}
			wg.Add(len(hooks))

			for _,hook:=range hooks{
				go func(h Hook){
					ctx,cancel:=context.WithTimeout(context.Background(),time.Second*30)
					err:=hook(ctx)
					if err!=nil{
						fmt.Printf("failed to run hook, err: %v \n", err)
					}
					cancel()
					wg.Done()
				}(hook)
			}
			wg.Wait()
			os.Exit(0)
	}
}