package main

import "fmt"

// 自定义类型
type MyInt int

// 类型别名
type AliasInt = int

// rune byte类型
// type byte = uint8

func main() {
	var i MyInt
	fmt.Println(i)
}
