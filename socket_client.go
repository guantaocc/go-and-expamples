package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	// 建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	defer conn.Close()
	if err != nil {
		fmt.Println("dial error")
	}
	for {
		input := bufio.NewReader(os.Stdin)
		s, err := input.ReadString('\n')
		if err != nil {
			fmt.Println("err")
			break
		}
		conn.Write([]byte(s))
	}
}
