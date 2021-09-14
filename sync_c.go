package main

import (
	"fmt"
	"sync"
)

// 并发安全和锁机制

var x int
var wg sync.WaitGroup
var lock sync.Mutex // 互斥锁
var rw sync.RWMutex // 读写锁

// map锁 sync.Map

// 并发的map
// load:加载  store: 存储
//var m sync.Map

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x++
		// 释放锁
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Print(x)
}
