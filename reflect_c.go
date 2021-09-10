package main

import (
	"fmt"
	"reflect"
)

// 反射

func reflectType(v interface{}) {
	// 反射获取结构体类型

	t := reflect.TypeOf(v)
	fmt.Printf("%v", t)
}

type Pairs struct {
	name string
}

func main() {
	var p = Pairs{}
	//reflectType(p)

	// 结构体类型解析
	// type name type kind

}
