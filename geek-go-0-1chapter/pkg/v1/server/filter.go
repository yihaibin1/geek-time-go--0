package server

import (
	"fmt"
	"time"
)

type handlerFunc func(c *Context)

type FilterBuild func(next Filter) Filter

type Filter func(c *Context)

var _ FilterBuild = MetricsFilterBuilder

func MetricsFilterBuilder(next Filter)Filter{
	return func(c *Context){
		start:=time.Now().Nanosecond()
		next(c)
		end:=time.Now().Nanosecond()
		fmt.Printf("转换filter花费了%d纳秒",end-start)
	}
}