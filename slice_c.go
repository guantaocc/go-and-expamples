package main

import "fmt"

func main() {
	a := []int{1, 2, 3}

	b := []int{2, 3}

	// 切片追加
	a = append(a, b...)
	fmt.Println(a)

	// 同一个切片copy
	a[4] = 100
	fmt.Println(b)

	c := make([]int, 5, 8)
	copy(c, a)
	fmt.Println(c)
}
