package main

import (
	"encoding/json"
	"fmt"
)

// 结构体可见性和json
// 大写开头则表示公开访问，小写则是私有的

// 结构体标签访问
type student struct {
	Id   int    `json:"Id"`
	Name string `json:"name"`
}

func newStudent(id int, name string) student {
	return student{
		Id:   id,
		Name: name,
	}
}

type class struct {
	Title    string    `json:"title"`
	Students []student `json:"students"`
}

func main() {
	// 班级
	c1 := class{
		Title:    "火箭",
		Students: make([]student, 0, 20),
	}
	// 往班级添加学生
	for i := 0; i < 10; i++ {
		c1.Students = append(c1.Students, newStudent(i, fmt.Sprintf("stu%02d", i)))
	}

	//fmt.Print(c1)
	// json序列化这个结构体数据
	data, err := json.Marshal(c1)
	if err != nil {

	}
	fmt.Printf("%T\n", data)
	//fmt.Printf("%s\n", data)

	// 反序列化
	var c2 class
	err = json.Unmarshal(data, &c2)
	if err != nil {

	}
	fmt.Printf("%v", c2)
}
