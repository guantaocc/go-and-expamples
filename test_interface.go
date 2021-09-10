package main

import "fmt"

// 测试接口的使用

type Sayer interface {
	say()
}

type Animal struct {
	name string
}

type Cat struct {
	Animal
	Type bool
}

type Dog struct {
	Animal
	Ohter string
}

func (c Cat) say() {
	fmt.Printf("%s is say", c.name)
}

func (d Dog) Write(p []byte) (int, error) {
	fmt.Print("writer")
	lenth := len(p)
	return lenth, nil
}

func (d Dog) say() {
	fmt.Printf("%s is say", d.name)
}

func da(s Sayer) {
	s.say()
}

func main() {
	var c = Cat{
		Animal: Animal{name: "cat"},
		Type:   false,
	}
	var d = Dog{
		Animal: Animal{name: "dog"},
		Ohter:  "",
	}
	da(c)
	da(d)
}
