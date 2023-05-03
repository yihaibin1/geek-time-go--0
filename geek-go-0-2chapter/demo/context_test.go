package demo

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContext(t *testing.T){
	ctx:=context.Background()
	timeCtx,cancel:=context.WithTimeout(ctx,time.Second*5)
	cancel()
	fmt.Println(timeCtx.Err())
}
