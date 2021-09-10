package main

// 空接口测试
import (
	"fmt"
)

func main() {
	// 类型断言
	var a interface{}
	a = 10
	a = map[string]string{"name": "1"}

	//_, ok := a.(string)
	//if !ok {
	//	fmt.Println("is not string")
	//}
	switch v := a.(type) {
	case string:
		fmt.Print(v)
	case int:
		fmt.Print(v)
	case map[string]string:
		fmt.Print(v)
	}
}
