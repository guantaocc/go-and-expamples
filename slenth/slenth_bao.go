package slenth

import "fmt"

// 包的使用和讲解

// 对外不可见
var name = "pony"

// 对外可见
var Name = "gucci"

func init() {
	// 会自动执行init
	fmt.Println("init is slenth bao")
}
