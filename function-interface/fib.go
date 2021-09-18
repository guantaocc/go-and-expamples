package main

import (
	"fmt"
	"io"
	"bufio"
	"strings"
)
// 利用fibnacie函数演示接口

type intGen func() int

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}


func (g intGen) Read(p []byte) (n int, err error){
	next := g()
	if next > 100 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	n, err = strings.NewReader(s).Read(p)
	fmt.Printf("p is %v\n", p)
	return
}

func printFileContents(reader io.Reader){
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main(){
	// var f intGen = Fibonacci()
	// printFileContents(f)
	s :=  "中文的的的SimpleNewReader returns a new Reader reading from s. " +
	"It is similar to bytes.NewBufferString but more efficient and read-only. "
	reader := strings.NewReader(s)
	var buf = make([]byte, 20)
	reader.Read(buf)
	fmt.Printf("%s", string(buf))
}