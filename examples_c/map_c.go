package main

import "fmt"

// 键值对测试

func main() {
	var a map[string]int

	fmt.Println(a == nil)

	// 初始化
	b := make(map[string]int, 8)
	fmt.Println(b)

	c := map[string]bool{
		"false": false,
		"狗子":    true,
	}
	fmt.Printf("%#v\n", c)

	// 判断键
	v, ok := c["狗子"]
	if ok {
		fmt.Println("存在", v)
	}

	// 遍历
	for k, v := range c {
		fmt.Println(k, v)
	}

	// 删除
	delete(c, "狗子")

}
