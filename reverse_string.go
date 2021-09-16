package main

import (
	"fmt"
	"strings"
)

func main() {
	// reverse string
	str := "hello, world"
	s := strings.Split(str, "")
	fmt.Printf("%v\n", s)
	left := 0
	right := len(s) - 1
	for left < right {
		s[left] = s[right]
		left++
		right--
	}
	fmt.Print(s)
}
