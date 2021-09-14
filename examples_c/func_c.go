package main

import "fmt"

// 函数的测试代码
func sum2num(a, b int) int {
	return a + b
}

// 函数值传递，字符，数字
func toName(name string) {
	name = "modify" + name
}

// 引用指针传递
func ptrName(name *string) *string {
	*name = "modify" + *name
	return name
}

// 可变参数
func sum3(arr ...int) (ret int) {
	for _, value := range arr {
		ret += value
	}
	return
}

func main() {
	// 函数声明，匿名函数，闭包
	sum2num(1, 2)
	name := "zhansan"

	ptrName(&name)

	fmt.Println(name)

	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(sum3(1, 2, 3, 4, 5))

	// 参数接片或数组展开赋值
	fmt.Println(sum3(arr...))
}
