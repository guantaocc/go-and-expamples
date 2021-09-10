package main

import "fmt"

// 操作字符串

func main() {
	s1 := "hello world liming worked in a resrurant 测试"
	s2 := "string sfs"
	// len, 一个中文为三个字节
	fmt.Print(len(s1))

	// 拼接
	str := fmt.Sprintf("%s - %s", s1, s2)
	fmt.Print(str)

}
