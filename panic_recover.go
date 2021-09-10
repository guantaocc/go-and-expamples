package main

import "fmt"

// panic操作

func a() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("a is error")
		}
	}()
	panic("aaa")
}

func b() {
	fmt.Println("bb")
}

func main() {
	a()
	b()
}
