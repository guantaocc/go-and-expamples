package main

import "fmt"

// 结构体
// 大写字符会暴露到包的外部

type people struct {
	name, city string
}

// 实例化结构体
var People people

// 构造函数
func NewPerson(name, city string) *people {
	return &people{
		name: name,
		city: city,
	}
}

// 定义方法
func (p *people) setCity(city string) {
	p.city = city
}

func main() {
	// 匿名结构体声明
	var user struct{ name, city string }
	user.name = "北京人是"
	user.city = "北京"
	fmt.Println(user)

	var p = new(people)
	p.name = "aa"

	// 指针实例化
	var p1 = &people{}
	p1.setCity("shanghai")
	fmt.Printf("%#v", p1)

	// 值列表初始化
	var p2 = &people{
		"闪动",
		"山东人",
	}
	fmt.Printf("%#v", p2)

}
