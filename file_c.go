package main

import (
	"bufio"
	"fmt"
	"os"
)

// 文件操作

func main() {
	// 读文件

	//f, err := os.Open("./tt.txt")
	//defer f.Close()
	//
	//if err != nil {
	//	fmt.Println(err)
	//}

	// 读io
	//for  {
	//	var temp  = make([]byte, 22)
	//
	//	n, err := f.Read(temp)
	//	if err == io.EOF {
	//		fmt.Println("end run ", string(temp[:n]))
	//		os.Exit(1)
	//	}
	//	if err != nil {
	//		fmt.Println("err read filed")
	//	}
	//	fmt.Println(string(temp))
	//}

	// 读 bufio
	//reader := bufio.NewReader(f)
	//for  {
	//	line, err := reader.ReadString('\n')
	//
	//	if err == io.EOF {
	//		fmt.Println(line)
	//		os.Exit(1)
	//	}
	//	if err != nil {
	//		fmt.Println("error")
	//	}
	//	fmt.Println(line)
	//}

	// ioutil ReadFile小文件读取

	// 写入文件
	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("error")
		return
	}
	defer file.Close()

	str := "小王子\n"
	//file.Write([]byte(str))
	//file.WriteString(str)

	// bufio
	writer := bufio.NewWriter(file)
	writer.WriteString(str)
	writer.Flush()
}
