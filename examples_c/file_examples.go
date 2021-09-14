package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// os.open读文件, 通常在读文件使用
func ReadByOsOpen() {
	fileObj, err := os.Open("./test.txt")
	defer fileObj.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 读取文件
	for {
		var tmp = make([]byte, 128)
		n, err := fileObj.Read(tmp)
		if err == io.EOF {
			fmt.Printf("%v", string(tmp[:n]))
			break
		}
		if err != nil {
			fmt.Println("读取失败")
		}
		fmt.Printf("%v", string(tmp))
	}
}

// 具有缓冲的type读取
func ReadByBufIo() {
	fileObj, err := os.Open("./test.txt")
	defer fileObj.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	rd := bufio.NewReader(fileObj)
	for {
		line, err := rd.ReadString('\n')
		if err == io.EOF {
			fmt.Printf("%v", line)
			break
		}
		fmt.Printf("%v", line)
	}
}

// 从ioutil中读取文件
func ReadByIoUtil() {
	contents, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		fmt.Printf("read file error %v", err)
		return
	}
	fmt.Print(string(contents))
}

func WriteByOs() {
	fileObj, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_APPEND|os.O_APPEND, 0644)
	if err != nil {
		fmt.Print("open filed err")
		return
	}
	str := "追加内容"
	//fileObj.Write([]byte(str))
	fileObj.WriteString(str)
}

// bufio
func WriteByBufio() {
	fileObj, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_APPEND|os.O_APPEND, 0644)
	if err != nil {
		fmt.Print("open filed err")
		return
	}
	writer := bufio.NewWriter(fileObj)
	str := "追加内容"
	writer.WriteString(str)
	// 将磁盘内容刷入
	writer.Flush()
}

// ioutil
func WriteByIoutil() {
	str := "write by io util"
	err := ioutil.WriteFile("./test.txt", []byte(str), 0644)
	if err != nil {
		print("write by ioitil err")
	}
}

func main() {
	//ReadByOsOpen()
	//ReadByBufIo()
	//ReadByIoUtil()
	//WriteByOs()
	//WriteByBufio()
	WriteByIoutil()
}
