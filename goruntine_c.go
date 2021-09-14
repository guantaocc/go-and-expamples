package main

import "sync"

// 并发编程
var wg sync.WaitGroup

func hello() {
	// 通知计数器-1
	defer wg.Done()
	print("world")
}

func main() {
	// 启动一个goruntine
	wg.Add(10000)

	for i := 0; i < 10000; i++ {
		go hello()
	}

	// main 的 goruntine 已经结束
	print("hello\n")
	// 等待计数器为0 则继续向下执行
	wg.Wait()
}
