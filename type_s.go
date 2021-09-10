package main

import (
	"fmt"
	"math"
)

// 数据类型

func main() {
	// 十进制
	var a int = 10
	fmt.Printf("%b\n", a)

	// 八进制
	b := 075
	fmt.Printf("%d\n", b)
	fmt.Printf("%o\n", b)

	// 十六进制
	c := 0xff
	fmt.Printf("%d\n", c)
	fmt.Printf("%x\n", c)

	// 浮点型
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)

	// bool默认false，不允许强制转换
	var d bool
	fmt.Printf("%s\n", d)

	// 字符 byte rune
	e := '中'
	fmt.Print(e)
}
