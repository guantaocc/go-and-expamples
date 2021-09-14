package main

import (
	"bufio"
	"fmt"
	"net"
)

// socket示例
func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("err")
			break
		}
		// 接收数据
		recv := string(buf[:n])
		fmt.Println(recv)
		// 返回数据
		conn.Write(buf[:n])
	}
}

func main() {
	// tcp通信
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Errorf("%s", err)
	}
	for {
		// 建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("%s", err)
			continue
		}
		go process(conn)
	}
}
