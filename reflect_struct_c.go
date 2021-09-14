package main

import (
	"fmt"
	"reflect"
)

// 结构体反射
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 首字母大小暴露方法
func (p Person) GetValidAge() int {
	return p.Age
}

func (p Person) GetValidName() string {
	return p.Name
}

func main() {
	// 解析结构体类型的反射

	// 初始化
	var p = Person{
		Name: "zhansan",
		Age:  0,
	}

	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	//fmt.Printf("%v, %v\n", t.Name(), t.Kind())
	//
	//// 遍历结构体字段
	//for i := 0; i < t.NumField(); i++ {
	//	filed := t.Field(i)
	//	fmt.Printf("name:%s, index:%d, tag: %v, type: %v\n", filed.Name, filed.Index, filed.Tag, filed.Type)
	//}
	//
	//// 名称反射字段
	//filed, ok := t.FieldByName("Name")
	//if ok {
	//	fmt.Printf("%v\n", filed)
	//}

	// 反射结构体的方法
	print(t.NumMethod())

	for i := 0; i < t.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		// name
		fmt.Printf("%v, %v\n", t.Method(i).Name, t.Method(i).Type)
		fmt.Printf("type:%s\n", methodType)

		// 调用方法
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}
