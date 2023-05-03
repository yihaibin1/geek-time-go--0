package test

import (
	"fmt"
	"reflect"
)

type users struct {
	name string
}

func main() {
	u := users{name: "alice"}
	fmt.Println(reflect.TypeOf(u))
	switch reflect.TypeOf(u) {
	case reflect.TypeOf(users{}):
		fmt.Println("我是 main.users")
	default:
		fmt.Println("不知道是啥")
	}
}
