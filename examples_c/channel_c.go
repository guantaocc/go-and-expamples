package main

var ch chan int

func run(c chan<- int) {
	// 只写值
	c <- 10
	// 关闭ch
	close(c)
}

func main() {

	// 初始化有缓冲区通道
	ch = make(chan int, 2)

	go run(ch)

	// 读取channel会阻塞
	for {
		// 是否还能够从通道读取值
		x, ok := <-ch
		print(x)
		if !ok {
			break
		}
	}
}
