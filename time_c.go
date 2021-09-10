package main

import (
	"fmt"
	"time"
)

// time库用法

func main() {
	now := time.Now()
	fmt.Printf("%s\n", now)

	// 秒数, 时间
	fmt.Println(now.Unix())
	fmt.Println(now.Year())
	fmt.Println(int(now.Month()))

	// 经过的事件
	time.Sleep(1 * time.Second)

	//for tep := range time.Tick(time.Second){
	//	fmt.Println(tep)
	//}

	// 时间格式化
	ret := now.Format("2006 01 02")
	fmt.Println(ret)

	// location
}
