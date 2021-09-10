package main

import "fmt"

// 结构体匿名字段
type Sitemap struct {
	string
	int
}

func main() {
	s := Sitemap{
		string: "小王子",
		int:    0,
	}
	fmt.Printf("%#v", s)
}
