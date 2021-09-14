package main

import "fmt"

// 结构体匿名字段
type Sitemap struct {
	string
	int
}

// 结构体嵌套和继承
type Shop struct {
	name    string
	address string
	saleDay int
	// 嵌套字段
	sales Salesman
}

type Salesman struct {
	green string
}

func main() {
	s := Sitemap{
		string: "小王子",
		int:    0,
	}
	fmt.Printf("%#v", s)

	// 嵌套结构体声明
	sh := Shop{
		name:    "aaa",
		address: "ff",
		saleDay: 0,
		sales: Salesman{
			green: "sss",
		},
	}
	fmt.Printf("%v", sh)
}
