package main

import (
	"fmt"
	"reflect"
)

// 普通类型反射

func reflectType(v interface{}) {
	// 反射获取结构体类型

	t := reflect.TypeOf(v)
	fmt.Printf("%v\n", t)
	// 不同的类型信息
	fmt.Printf("%s, %s\n", t.Name(), t.Kind())
}

// 值的类型
func reflectValue(v interface{}) {
	t := reflect.ValueOf(v)
	fmt.Printf("%s\n", t)
}

func reflectSetValue(v interface{}) {
	t := reflect.ValueOf(v)
	k := t.Elem().Kind()
	switch k {
	case reflect.Int32:
		// 反射修改指针类型的值
		t.Elem().SetInt(32)
	}
}

type Pairs struct {
	name string
}

func main() {
	var p = Pairs{}
	reflectType(p)

	reflectValue(p)

	// 指针类型的反射修改
	var a int32 = 120
	reflectSetValue(&a)
	fmt.Printf("%v\n", a)
}
